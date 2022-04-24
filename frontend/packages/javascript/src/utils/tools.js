export default {
    createAvatarByName(name) {
        let word = name.substring(1, 0);
        let str = '';
        for (var i = 0; i < name.length; i++) {
            str += parseInt(name[i].charCodeAt(0), 10).toString(16);
        }
        return {
            backgroundColor: `#${str}`,
            word
        }
    }
}