
let modules = {};
const files = import.meta.globEager('./*.js');
for (const key in files) {
  const fileName = key.replace(/^\.\/(.*)\.\w+$/, '$1')
  modules[fileName] = files[key].default || files[key]
}

export default {
  ...modules
}