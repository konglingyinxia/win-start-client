export namespace api {
	
	export class AppManager {
	    separator: string;
	    appHomeKey: string;
	    appHomePath: string;
	
	    static createFrom(source: any = {}) {
	        return new AppManager(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.separator = source["separator"];
	        this.appHomeKey = source["appHomeKey"];
	        this.appHomePath = source["appHomePath"];
	    }
	}
	export class EnvManager {
	    separator: string;
	    envHomeKey: string;
	    envHomePath: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvManager(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.separator = source["separator"];
	        this.envHomeKey = source["envHomeKey"];
	        this.envHomePath = source["envHomePath"];
	    }
	}

}

export namespace dto {
	
	export class DiskCount {
	    total: number;
	    free: number;
	    used: number;
	    usedPercent: number;
	
	    static createFrom(source: any = {}) {
	        return new DiskCount(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.free = source["free"];
	        this.used = source["used"];
	        this.usedPercent = source["usedPercent"];
	    }
	}
	export class HomeOsInfo {
	    infoStat: host.InfoStat;
	    ipAddr: string;
	
	    static createFrom(source: any = {}) {
	        return new HomeOsInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.infoStat = this.convertValues(source["infoStat"], host.InfoStat);
	        this.ipAddr = source["ipAddr"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SystemInfo {
	    infoStat: host.InfoStat;
	    ipAddr: string;
	    rootDir: string;
	    cpu: string;
	    gpu: string;
	    memory: any;
	    chassis: string;
	    disk: DiskCount;
	    network: string;
	    product: string;
	    baseboard: string;
	    bios: string;
	    manufacturer: string;
	    productName: string;
	
	    static createFrom(source: any = {}) {
	        return new SystemInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.infoStat = this.convertValues(source["infoStat"], host.InfoStat);
	        this.ipAddr = source["ipAddr"];
	        this.rootDir = source["rootDir"];
	        this.cpu = source["cpu"];
	        this.gpu = source["gpu"];
	        this.memory = source["memory"];
	        this.chassis = source["chassis"];
	        this.disk = this.convertValues(source["disk"], DiskCount);
	        this.network = source["network"];
	        this.product = source["product"];
	        this.baseboard = source["baseboard"];
	        this.bios = source["bios"];
	        this.manufacturer = source["manufacturer"];
	        this.productName = source["productName"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace host {
	
	export class InfoStat {
	    hostname: string;
	    uptime: number;
	    bootTime: number;
	    procs: number;
	    os: string;
	    platform: string;
	    platformFamily: string;
	    platformVersion: string;
	    kernelVersion: string;
	    kernelArch: string;
	    virtualizationSystem: string;
	    virtualizationRole: string;
	    hostId: string;
	
	    static createFrom(source: any = {}) {
	        return new InfoStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.uptime = source["uptime"];
	        this.bootTime = source["bootTime"];
	        this.procs = source["procs"];
	        this.os = source["os"];
	        this.platform = source["platform"];
	        this.platformFamily = source["platformFamily"];
	        this.platformVersion = source["platformVersion"];
	        this.kernelVersion = source["kernelVersion"];
	        this.kernelArch = source["kernelArch"];
	        this.virtualizationSystem = source["virtualizationSystem"];
	        this.virtualizationRole = source["virtualizationRole"];
	        this.hostId = source["hostId"];
	    }
	}

}

export namespace model {
	
	export class FormatterTime {
	
	
	    static createFrom(source: any = {}) {
	        return new FormatterTime(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class GlobalSetting {
	    id: string;
	    // Go type: FormatterTime
	    createdAt: any;
	    // Go type: FormatterTime
	    updatedAt: any;
	    defaultWeb: string;
	    openDefaultWeb?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GlobalSetting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.defaultWeb = source["defaultWeb"];
	        this.openDefaultWeb = source["openDefaultWeb"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace req {
	
	export class AppReq {
	    id: string;
	    // Go type: model
	    createdAt: any;
	    // Go type: model
	    updatedAt: any;
	    type: string;
	    name: string;
	    version: string;
	    appDir: string;
	    envId: string;
	    startCmd: string;
	    stopCmd: string;
	    restartCmd: string;
	    versionCmd: string;
	    logDir: string;
	    username: string;
	    password: string;
	    host: string;
	    port: number;
	    initDb: boolean;
	    autoStart: boolean;
	    startDelay: number;
	    startOrder: number;
	    envVars: string;
	    remark: string;
	
	    static createFrom(source: any = {}) {
	        return new AppReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.type = source["type"];
	        this.name = source["name"];
	        this.version = source["version"];
	        this.appDir = source["appDir"];
	        this.envId = source["envId"];
	        this.startCmd = source["startCmd"];
	        this.stopCmd = source["stopCmd"];
	        this.restartCmd = source["restartCmd"];
	        this.versionCmd = source["versionCmd"];
	        this.logDir = source["logDir"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.initDb = source["initDb"];
	        this.autoStart = source["autoStart"];
	        this.startDelay = source["startDelay"];
	        this.startOrder = source["startOrder"];
	        this.envVars = source["envVars"];
	        this.remark = source["remark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class EnvCheckReq {
	    id: string;
	    command: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvCheckReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.command = source["command"];
	    }
	}
	export class EnvReqAdd {
	    id: string;
	    envName: string;
	    envDir: string;
	    envVars: string;
	    remark: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvReqAdd(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.envName = source["envName"];
	        this.envDir = source["envDir"];
	        this.envVars = source["envVars"];
	        this.remark = source["remark"];
	    }
	}

}

export namespace res {
	
	export class BaseRes {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new BaseRes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}
	export class OverviewData {
	    runningNum: number;
	    stoppedNum: number;
	    envNum: number;
	    storageApp: number;
	    storageEnv: number;
	    storageLog: number;
	
	    static createFrom(source: any = {}) {
	        return new OverviewData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.runningNum = source["runningNum"];
	        this.stoppedNum = source["stoppedNum"];
	        this.envNum = source["envNum"];
	        this.storageApp = source["storageApp"];
	        this.storageEnv = source["storageEnv"];
	        this.storageLog = source["storageLog"];
	    }
	}

}

