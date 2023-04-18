export function computeWinterEquipmentTotal(input) {
    let sum = 0;
    let current;
    for (current of input) sum += current;
    return sum;
}
