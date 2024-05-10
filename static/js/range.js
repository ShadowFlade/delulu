document.addEventListener("DOMContentLoaded", function() {
    const rangeEl = document.querySelector(".js-calculator__range-input");
    if (!rangeEl) return;
    const rangeLabelEl = document.querySelector(".js-range-input__label");
    if (!rangeLabelEl) return;
    const ratio = rangeLabelEl.offsetWidth / rangeEl.offsetWidth;

    rangeLabelEl.dataset.ratio = ratio.toFixed(2);

    rangeEl.addEventListener("input", function() {
        rangeLabelEl.textContent = rangeEl.value;
        const translate = calcTooltipTranslate(rangeEl, rangeLabelEl);
    });
});

function calcTooltipTranslate(rangeEl, rangeLabelEl) {
    let translate;
    const valueDiff = +rangeEl.max - +rangeEl.min;

    translate =
        ((+(+rangeEl.value) - +rangeEl.min) / valueDiff) *
            (parseFloat(rangeEl.offsetWidth) -
                parseFloat(rangeLabelEl.offsetWidth) / 2) - 15;

    //this is kinda terrible tbh
    const resultString = `translateX(${translate}px)`;
    rangeLabelEl.style.transform = resultString;
}
