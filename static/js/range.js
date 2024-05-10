document.addEventListener("DOMContentLoaded", function() {
    const rangeEl = document.querySelector(".js-calculator__range-input");
    if (!rangeEl) return;
    const rangeLabelEl = document.querySelector(".js-range-input__label");
    if (!rangeLabelEl) return;
    const ratio = rangeLabelEl.offsetWidth / rangeEl.offsetWidth;

    rangeLabelEl.dataset.ratio = ratio.toFixed(2);
    rangeEl.addEventListener("input", function() {
        rangeLabelEl.textContent = rangeEl.value;
        let translate;
        let correction;
        if (rangeEl.value != rangeEl.min) {
            const valueDiff = +rangeEl.max - +rangeEl.min;
            const isPositiveCorrection =
                +rangeEl.value - +rangeEl.min > (+rangeEl.max - +rangeEl.min) / 2;
            translate =
                ((+rangeEl.value - +rangeEl.min) / (+rangeEl.max - +rangeEl.min) || 1) *
                (parseFloat(rangeEl.offsetWidth) -
                    parseFloat(rangeLabelEl.offsetWidth)) +
                (isPositiveCorrection
                    ? +rangeLabelEl.offsetWidth / 2
                    : -+rangeLabelEl.offsetWidth / 2);

            //this is kinda terrible tbh
            correction = isPositiveCorrection
                ? -+rangeEl.value / 12
                : +rangeEl.value / 12;
        } else {
            translate = 0;
        }
        const resultString = `translateX(${translate + correction}px)`;
        rangeLabelEl.style.transform = resultString;
    });
});
