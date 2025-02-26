//写个函数,根据字节大小转化为GB，MB，KB
export function formatBytes(bytes: number) {
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    let i = 0;
    while (bytes >= 1024) {
        bytes /= 1024;
        i++;
    }
    return {size: bytes.toFixed(2), unit: sizes[i]};
}

//写个函数时间戳格式化
export function formatDate(timestamp: number) {
    const date = new Date(timestamp);
    return date.toLocaleString();
}

//把秒时间戳转化为  3天 4小时 5分钟 6秒
export function formatDateDHMS(newSeconds: number) {
    const d = Math.floor(newSeconds / (24 * 60 * 60));
    const h = Math.floor(newSeconds % (24 * 60 * 60) / 3600);
    const m = Math.floor((newSeconds % (24 * 60 * 60) % 3600) / 60);
    const s = newSeconds % 60;
    return {day: d, hours: h, minutes: m, seconds: s};
}
