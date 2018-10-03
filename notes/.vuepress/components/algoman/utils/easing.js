function easeInBack(ratio) {
    let s = 1.70158
    return ratio * ratio * ((s + 1.0) * ratio - s);
}

function easeOutBack(ratio) {
    let s = 1.70158;
    ratio = ratio - 1.0;
    return ratio * ratio * ((s + 1.0) * ratio + s) + 1.0;
}

export default function easeInOutBack(ratio) {
    ratio = ratio * 4.0;
    return (ratio < 1.0) ? 0.5 * easeInBack(ratio) : 0.5 * easeOutBack(ratio - 1.0) + 0.5;
}

export function powerEasing(ratio) {
    let x = 0.5
    return Math.pow(ratio, 2) * ((x+1)*ratio - x)
}