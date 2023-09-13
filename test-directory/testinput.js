mport { test } from '@/utils/test.js';
import { other } from '@/stuff/thing/other.js';

let arr = [];

for (let i = 0; i < 100; i++) {
    arr[i] = test();
    if (other.thing) {
        arr[i] += 100;
    }
}

