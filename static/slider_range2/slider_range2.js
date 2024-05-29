document.addEventListener("DOMContentLoaded", () => {
	const slider = document.querySelector('#js-slider-range-age');
	console.log(slider,' slider')
	const minHiddenInput = document.querySelector("[name='age-min']")
	const maxHiddenInput = document.querySelector("[name='age-max']")
	if(!minHiddenInput || !maxHiddenInput) {
		return;
	}
	//TODO add debouncer
	slider.addEventListener('change', (evt) => {
		minHiddenInput.value = evt.detail.value1;
		maxHiddenInput.value = evt.detail.value2;
	});
})
