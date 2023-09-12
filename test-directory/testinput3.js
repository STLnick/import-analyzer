import { test } from '@/utils/testBar';
import { other } from '@/stuff/thing/otherFoo';
import Thing from '@/classes/thingTwo';
import ChildThing from '@/classes/nested/childThing';

let arr = [];

for (let i = 0; i < 100; i++) {
    arr[i] = test();
    if (other.thing && Thing.staticMethod() && ChildThing.staticMethod()) {
        arr[i] += 100;
    }
}

