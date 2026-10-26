/**
 * @param {string} title
 * @return {string}
 */
var capitalizeTitle = function(title) {
    const chars = title.toLowerCase().split('');
    const n = chars.length;
    let i = 0;
    
    while(i < n){
        let j = i;
        while(j < n && chars[j] !== ' '){
            j++;
        }

        if (j - i > 2){
            chars[i] = chars[i].toUpperCase();
        }

        i = j + 1;
    }

    return chars.join('');
};