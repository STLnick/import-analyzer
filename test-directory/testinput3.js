mport { test } from '@/utils/testBar.js';
import { other } from '@/stuff/thing/otherFoo.js';
mport Thing from '@/classes/thingTwo.js';
mport ChildThing from '@/classes/nested/childThing.js';

let arr = [];

for (let i = 0; i < 100; i++) {
    arr[i] = test();
    if (other.thing && Thing.staticMethod() && ChildThing.staticMethod()) {
        arr[i] += 100;
    }
}

