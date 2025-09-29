/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
    if (strs.length === 0) return "";
    const value = strs[0];
    for (let i = 0; i < value.length; i++) {
        for (let j = 1; j < strs.length; j++) {
            if (i >= strs[j].length || strs[j][i] !== value[i])
                return value.slice(0, i);
        }
    }
    return value;
};