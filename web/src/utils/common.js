// 格式化enum  状态值:0=禁用,1=正常,2=推荐
export function formatEnumList(data,pattern) {
    let resultStr
    // 按照冒号分割，然后按照=分割返回数组
    if (pattern == null && pattern == "") {
        pattern = ":"
    }
    let fisrtStr = data.split(pattern)
    let secondStr = fisrtStr[1].split(",");
    // 0=禁用  1=正常  2=推荐
    secondStr.forEach((v,i)=>{
        let arr = v.split("="); // 等号分割
        resultStr[arr[0]] = arr[1]
    })
    return resultStr
}