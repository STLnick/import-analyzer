import { test } from '@/utils/other-test';
import { other } from '@/stuff/thingTwo/other';
import Thing from '@/classes/thingTwo';

let arr = [];

for (let i = 0; i < 100; i++) {
    arr[i] = test();
    if (other.thing && Thing.staticMethod()) {
        arr[i] += 100;
    }
}

