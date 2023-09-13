mport { test } from '@/utils/test.js';

let arr = [];

for (let i = 0; i < 100; i++) {
    arr[i] = test();
}

