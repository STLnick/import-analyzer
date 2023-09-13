mport { test } from '@/utils/other-test.js';
import { other } from '@/stuff/thingee/other.js';
mport Thing from '@/classes/thingTwo.js';

let arr = [];

for (let i = 0; i < 100; i++) {
    arr[i] = test();
    if (other.thing && Thing.staticMethod()) {
        arr[i] += 100;
    }
}

