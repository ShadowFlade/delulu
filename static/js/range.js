document.addEventListener("DOMContentLoaded", function() {
    initSliders();
});

function calcTooltipTranslate(rangeEl, rangeLabelEl) {
    let translate;
    const valueDiff = +rangeEl.max - +rangeEl.min;

    translate =
        ((+(+rangeEl.value) - +rangeEl.min) / valueDiff) *
            (parseFloat(rangeEl.offsetWidth) -
                parseFloat(rangeLabelEl.offsetWidth) / 2) -
            15;

    //this is kinda terrible tbh
    const resultString = `translateX(${translate}px)`;
    if (rangeEl.dataset.type == "money") {
        const newValue = formatValueToMoney(rangeEl.value, +rangeEl.max, +rangeEl.min);
        rangeLabelEl.textContent = newValue;
    }
     else {
        rangeLabelEl.textContent = rangeEl.value;
    }
    rangeLabelEl.style.transform = resultString;
}

function initSliders() {
    const rangeEl = document.querySelectorAll(".js-calculator__range-input");
    if (!rangeEl) return;
    const rangeLabelElAll = document.querySelectorAll(".js-range-input__label");
    if (!rangeLabelElAll) return;
    rangeEl.forEach((rangeEl, index) => {
        const rangeLabelEl = rangeLabelElAll[index];
        calcTooltipTranslate(rangeEl, rangeLabelEl);
        const ratio = rangeLabelEl.offsetWidth / rangeEl.offsetWidth;

        rangeLabelEl.dataset.ratio = ratio.toFixed(2);

        rangeEl.addEventListener("input", function() {
            const translate = calcTooltipTranslate(rangeEl, rangeLabelEl);
        });
    });
}

function formatValueToMoney(value, max, min) {
    const money = value;
    const formatter = new Intl.NumberFormat("ru-ru", {
        notation: "compact",
        minimumFractionDigits: 0,
        maximumSignificantDigits: 3,
        style: "currency",
        currency: "RUB",
    });

    let newValue = formatter.format(money);

    if (value == max) {
        newValue = '>' + newValue;
    } else if (value == min){
        newValue = "Любой"
    }

    return newValue;
}
