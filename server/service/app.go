package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/duke-git/lancet/v2/system"
	"github.com/jinzhu/copier"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto/req"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/init/log"
	"github.com/konglingyinxia/win-start-client/server/model"
	"github.com/konglingyinxia/win-start-client/server/utils/env"
	"github.com/konglingyinxia/win-start-client/server/utils/id"
	"github.com/konglingyinxia/win-start-client/server/utils/osutils"
	"github.com/shirou/gopsutil/v4/process"
	runtime2 "github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type AppService struct{}

type IAppService interface {
	Add(req req.AppReq) error
	List() (*[]res.AppRes, error)
	ListAll() (apps []model.App, err error)
	ListByIds(ids []string) (apps []*model.App, err error)
	Delete(id string) error
	CountByEnvID(id string) (int64, error)
	GetById(id string) (*model.App, error)
	Update(appReq req.AppReq) error
	RunningStatus(apps []model.App) ([]res.AppRunStatus, error)
	Start(app *model.App) error
	Stop(app *model.App) error
	Restart(app *model.App) error
	StartAll(isInitial bool, apps []*model.App) error
	StopAll(isShutdown bool, apps []*model.App) error
	RefreshPorts(app model.App) (ports []int, err error)
}

func NewAppService() IAppService {
	return &AppService{}
}

func (a AppService) Add(req req.AppReq) error {
	app := model.App{}
	err := copier.Copy(&app, req)
	if err != nil {
		return err
	}
	app.ID = id.UuidSimple()
	//日志路径
	logPath := constant.GetAppLogRelativePath(app.ID, app.Name)
	app.LogDir = logPath
	err = global.DB.Create(&app).Error
	return err
}
func (a AppService) List() (*[]res.AppRes, error) {
	var apps []model.App
	err := global.DB.Order("start_order asc,created_at desc").Find(&apps).Error
	if err != nil {
		return nil, err
	}
	var appRes []res.AppRes
	for _, app := range apps {
		t := res.AppRes{}
		err := copier.Copy(&t, &app)
		if err != nil {
			return nil, err
		}
		//应用全路径
		fullPath := constant.GetApplicationFullPath(app.AppDir)
		diskSize := osutils.DirSize(fullPath)
		t.DiskSize = diskSize
		appRes = append(appRes, t)
	}
	return &appRes, err
}

func (a AppService) ListAll() (apps []model.App, err error) {
	err = global.DB.Find(&apps).Order("created_at desc").Error
	return
}

func (a AppService) ListByIds(ids []string) (apps []*model.App, err error) {
	apps = []*model.App{}
	err = global.DB.Find(&apps, "id in (?)", ids).Error
	return apps, err
}

func (a AppService) Delete(id string) error {
	err := global.DB.Delete(&model.App{}, "id = ?", id).Error
	return err
}

func (a AppService) CountByEnvID(envId string) (int64, error) {
	var count int64
	//查询数量
	err := global.DB.Model(&model.App{}).Where("env_id = ?", envId).Count(&count).Error
	return count, err
}
func (a AppService) GetById(id string) (*model.App, error) {
	var app model.App
	err := global.DB.Where("id = ?", id).First(&app).Error
	return &app, err
}
func (a AppService) Update(appReq req.AppReq) error {
	var app model.App
	err := global.DB.Where("id = ?", appReq.ID).First(&app).Error
	if err != nil {
		return err
	}
	err = copier.Copy(&app, appReq)
	if err != nil {
		return err
	}
	err = global.DB.Model(&model.App{}).Select("*").Where("id = ?", appReq.ID).Updates(app).Error
	return err
}
func (a AppService) RunningStatus(apps []model.App) (runStatus []res.AppRunStatus, err error) {
	var result []res.AppRunStatus
	for _, app := range apps {
		runStatus := res.AppRunStatus{}
		//进程状态
		procPath := constant.GetProcessFullPath(app.ID)
		pid := -1
		if fileutil.IsExist(procPath) {
			pidStr, err := fileutil.ReadFileToString(procPath)
			if err == nil {
				pid, _ = strconv.Atoi(pidStr)
			}
		}
		runStatus.Status = constant.RuntimeStopped
		if pid > -1 {
			exists, err := process.PidExists(int32(pid))
			if err != nil {
				continue
			}
			if exists {
				proc, err := process.NewProcess(int32(pid))
				if err != nil {
					continue
				}
				if system.IsWindows() {
					runStatus.Status = constant.RuntimeRunning
				} else {
					status, err := proc.Status()
					if err == nil {
						if slice.Contain(status, process.Stop) || slice.Contain(status, process.Zombie) {
							runStatus.Status = constant.RuntimeStopped
						} else if slice.Contain(status, process.Idle) {
							runStatus.Status = constant.RuntimePadding
						} else {
							runStatus.Status = constant.RuntimeRunning
						}
					}
				}
				runStatus.CpuPercent, err = proc.CPUPercent()
				runStatus.CpuTime, err = proc.Times()
				runStatus.MemPercent, err = proc.MemoryPercent()
				runStatus.Memory, err = proc.MemoryInfo()
				if runStatus.Memory != nil {
					runStatus.MemSize += runStatus.Memory.RSS
				}
				runStatus.Pid = int(proc.Pid)
				//获取子进程
				children, err := osutils.ProcChildren(proc)
				if err != nil {
					global.LOG.Errorf("[%s]子进程[%d]信息获取失败：%v", app.Name, pid, err)
				}
				runStatus.Pids = append(runStatus.Pids, pid)
				if children != nil && len(children) > 0 {
					for _, child := range children {
						cpuPerc, err := child.CPUPercent()
						if err == nil {
							runStatus.CpuPercent += cpuPerc
						}
						memPerc, err := child.MemoryPercent()
						if err == nil {
							runStatus.MemPercent += memPerc
						}
						memInfo, err := child.MemoryInfo()
						if err == nil {
							runStatus.MemSize += memInfo.RSS
						}
						runStatus.Pids = append(runStatus.Pids, int(child.Pid))
					}
				}
			}
		}
		result = append(result, runStatus)
	}
	return result, nil
}

func (a AppService) Start(app *model.App) error {
	status, err := a.RunningStatus([]model.App{*app})
	if err != nil {
		return err
	}
	runStatus := status[0]
	if runStatus.Status != constant.RuntimeStopped {
		return errors.New("应用已经启动")
	}
	procPath := constant.GetProcessFullPath(app.ID)
	//启动应用
	newPid, err := startApp(app)
	if err != nil {
		return err
	}
	//更新进程文件
	err = fileutil.WriteStringToFile(procPath, strconv.Itoa(int(newPid)), false)
	if err != nil {
		return err
	}
	return nil
}

func startApp(app *model.App) (int32, error) {
	var pid int32
	cmd := assembleCmd(app, app.StartCmd)
	//启动应用
	if err := cmd.Start(); err != nil {
		global.LOG.Error(fmt.Sprintf("启动应用失败：%s", err.Error()))
		return pid, err
	}
	pid = int32(cmd.Process.Pid)
	go updateAppRunPort(app, []int32{pid}, true)
	return pid, nil
}

// 组装
func assembleCmd(app *model.App, execCmd string) *exec.Cmd {
	//空格分割Command
	commands := strutil.SplitAndTrim(execCmd, " ")
	//解析环境变量/
	resultEnv := env.ParseAllEnv(parseMergeEnv(app))
	for key, val := range resultEnv {
		os.Setenv(key, val)
	}
	//执行命令
	var cmd *exec.Cmd
	if len(commands) == 1 {
		cmd = exec.Command(commands[0])
	} else {
		cmd = exec.Command(commands[0], commands[1:]...)
	}
	//设置cmd不同平台参数
	osutils.CmdAddOsParams(cmd)
	//设置执行目录
	cmd.Dir = constant.GetApplicationFullPath(app.AppDir)
	//设置日志
	lumberlog := log.NewLumberlog(constant.GetAppLogRelativePath(app.ID, app.Name))
	cmd.Stdout = lumberlog
	cmd.Stderr = lumberlog
	return cmd
}

// 进程启动后，等待一定时间获取端口
// param app *model.App 应用对象
// param pid int 进程ID
// param auto bool true-自动获取 false-手动获取
func updateAppRunPort(app *model.App, pids []int32, auto bool) (ports []int) {
	for _, item := range []int{1, 2, 3} {
		// 自动获取端口需要等待一定时间
		if auto {
			time.Sleep(time.Second * 10 * time.Duration(item))
		}
		for _, pid := range pids {
			ports = append(ports, ProcessService.ProcessPortsByPid(pid)...)
		}
		ports = slice.Unique(ports)
		if len(ports) > 0 {
			jsonStr, _ := json.Marshal(ports)
			global.DB.Model(&model.App{}).Where("id = ?", app.ID).Updates(&model.App{Ports: string(jsonStr)})
		}
		// 手动获取端口，不需要等待
		if !auto {
			break
		}
	}
	return ports
}

func parseMergeEnv(app *model.App) map[string]string {
	//内置环境变量
	var envMaps map[string]string
	if app.EnvVars != "" {
		_ = json.Unmarshal([]byte(app.EnvVars), &envMaps)
	}
	//外部引用环境变量
	var envs map[string]string
	if app.EnvId != "" {
		e := model.Env{}
		global.DB.First(&e, "id = ?", app.EnvId)
		if e.EnvVars != "" {
			_ = json.Unmarshal([]byte(e.EnvVars), &envs)
		}
	}
	envMaps = maputil.Merge(envMaps, envs)
	return envMaps
}
func (a AppService) Stop(app *model.App) error {
	//进程状态
	procPath := constant.GetProcessFullPath(app.ID)
	if !fileutil.IsExist(procPath) {
		return errors.New("应用未启动")
	}
	pid := -1
	pidStr, err := fileutil.ReadFileToString(procPath)
	if err != nil {
		return err
	}
	pid, err = strconv.Atoi(pidStr)
	exists, err := process.PidExists(int32(pid))
	if err != nil {
		return err
	}
	if exists {
		//停止命令
		if app.StopCmd != "" {
			cmd := assembleCmd(app, app.StopCmd)
			//终止应用
			if err := cmd.Run(); err != nil {
				global.LOG.Error(fmt.Sprintf("终止应用【%s】失败：%s", app.Name, err.Error()))
				return err
			}
			//通知进程停止
		} else {
			if system.IsWindows() {
				pids, err := osutils.ProcGroupPids(uint32(pid))
				pids = append(pids, uint32(pid))
				err = osutils.KillProcess(pids)
				if err != nil {
					global.LOG.Error(fmt.Sprintf("终止应用【%s】失败：%s", app.Name, err))
					return err
				}
			} else {
				proc, err := process.NewProcess(int32(pid))
				if err != nil {
					return err
				}
				err = proc.Suspend()
				if err != nil {
					global.LOG.Error(fmt.Sprintf("暂停应用【%s】失败：%s", app.Name, err.Error()))
				}
				time.Sleep(time.Second * 2)
				err = proc.Terminate()
				if err != nil {
					global.LOG.Error(fmt.Sprintf("终止应用【%s】失败：%s", app.Name, err.Error()))
				}
				time.Sleep(time.Second * 2)
				err = proc.Kill()
				if err != nil {
					global.LOG.Error(fmt.Sprintf("杀死应用【%s】失败：%s", app.Name, err.Error()))
				}
			}
		}
	}
	//删除进程文件
	err = os.Remove(procPath)
	if err != nil {
		return err
	}
	return nil
}

func (a AppService) Restart(app *model.App) error {
	if app.RestartCmd != "" {
		cmd := assembleCmd(app, app.RestartCmd)
		//终止应用
		if err := cmd.Run(); err != nil {
			global.LOG.Error(fmt.Sprintf("重启应用【%s】失败：%s", app.Name, err.Error()))
			return err
		}
		return nil
	} else {
		_ = a.Stop(app)
		return a.Start(app)
	}
}

func (a AppService) StartAll(isInitial bool, apps []*model.App) (err error) {
	if apps == nil {
		err := global.DB.Order("start_order asc").Find(&apps).Error
		if err != nil {
			return err
		}
	}
	for _, app := range apps {
		//是初始启动//并且是自启动
		if isInitial {
			if *app.AutoStart {
				send := res.AppStartEvent{
					BaseModel: model.BaseModel{
						ID: app.ID,
					},
					Status:      "started",
					EventStatus: "running",
					EventType:   "start",
				}
				runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName, send)
				global.LOG.Info(fmt.Sprintf("应用自启-应用【%s】开始启动....", app.Name))
				time.Sleep(time.Second * time.Duration(app.StartDelay))
				err := a.Start(app)
				if err != nil {
					send.Status = "failed"
					send.Error = err.Error()
					runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName, send)
					global.LOG.Error(fmt.Sprintf("应用自启-启动应用【%s】失败：%s", app.Name, err.Error()))
				} else {
					send.Status = "running"
					runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName, send)
					global.LOG.Info(fmt.Sprintf("应用自启-应用【%s】启动成功", app.Name))
				}
			}
		} else {
			global.LOG.Info(fmt.Sprintf("应用启动ALL-应用【%s】开始启动....", app.Name))
			time.Sleep(time.Second * time.Duration(app.StartDelay))
			err := a.Start(app)
			if err != nil {
				global.LOG.Error(fmt.Sprintf("应用启动ALL-应用【%s】失败：%s", app.Name, err.Error()))
			} else {
				global.LOG.Info(fmt.Sprintf("应用启动ALL-应用【%s】启动成功", app.Name))
			}
		}
	}
	return err
}

func (a AppService) StopAll(isShutdown bool, apps []*model.App) (err error) {
	if apps == nil {
		err := global.DB.Order("start_order desc").Find(&apps).Error
		if err != nil {
			return err
		}
	}
	for _, app := range apps {
		send := res.AppStartEvent{
			BaseModel: model.BaseModel{
				ID: app.ID,
			},
			Status:      "stopping",
			EventStatus: "running",
			EventType:   "stop",
		}
		if isShutdown {
			runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName, send)
		}
		global.LOG.Info(fmt.Sprintf("应用停止-应用【%s】开始关闭....", app.Name))
		time.Sleep(time.Second * 1)
		err := a.Stop(app)
		if err != nil {
			if isShutdown {
				send.Status = "failed"
				send.Error = err.Error()
				runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName, send)
			}
			global.LOG.Error(fmt.Sprintf("应用停止-关闭应用【%s】失败：%s", app.Name, err.Error()))
		} else {
			if isShutdown {
				send.Status = "stopped"
				runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName, send)
			}
			global.LOG.Info(fmt.Sprintf("应用停止-应用【%s】关闭成功", app.Name))
		}
	}
	return err
}

func (a AppService) RefreshPorts(app model.App) (ports []int, err error) {
	allStatus, err := a.RunningStatus([]model.App{app})
	if err != nil {
		return nil, err
	}
	status := allStatus[0]
	if status.Status == constant.RuntimeStopped {
		return nil, errors.New("应用未启动")
	}
	result := slice.FlatMap(status.Pids, func(index int, item int) []int32 {
		return []int32{int32(item)}
	})
	ports = updateAppRunPort(&app, result, false)
	if len(ports) == 0 {
		return nil, errors.New("应用端口获取失败")
	}
	return ports, nil
}
