<template>
	<div>
		<main>
			<section>
				<div>
					<img
						src="https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcTl0E7Bc-YAWQtr0n5XVYBXNeUmsNM9MdchIOhnheuB1GgQGIlo"
						alt="性別&年齡"
						style="margin-bottom: 200px; width: 200px"
					/>
				</div>
				<div>
					<img
						src="https://steam.oxxostudio.tw/webp/python/example/matplotlib-bar-01.webp"
						alt="違規原因"
						style="width: 200px"
					/>
				</div>
			</section>
			<section>
				<div class="circle-chart">
					<div class="center-circle">違規事件<br />分析</div>
					<div class="section top-left">性別<br />&<br />年齡</div>
					<div class="section top-right">地區</div>
					<div class="section bottom-left">違規原因</div>
					<div class="section bottom-right">車種</div>
					<div class="circle-lines"></div>
				</div>
			</section>
			<section>
				<div>
					<img
						src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhq29lnk6xcqHxNssWvqH8SZa7GVmZ1iCva4iviB_5zmvtbka8_1m-uuXnAPxFyST4JMwc3AQXNNgILh0g2Fxff7fmWVcEoFKwt1JWKO0i0ZvOs66ili3XNxaARUgKkTJZmtYsfN8Bb2Ws/s1600/1.jpg"
						alt="地區"
						style="height: 200px; margin-bottom: 100px"
					/>
				</div>
				<div>
					<img
						src="https://lh3.googleusercontent.com/ZRbzormInrqWNrjKjb21Rolv41Pm7psNuEOdkceeTRhZZljB-Lxgx6d3QA9oe1V-Kwus=w280"
						alt="車種"
						style="margin-left: -50px"
					/>
				</div>
			</section>
		</main>
		<div>
			<button
				style="
					font-size: 20px;
					position: fixed;
					left: 500px;
					bottom: 100px;
				"
			>
				返回
			</button>
		</div>
	</div>
</template>

<script>
export default {
	name: "ViolationManagement",
};
</script>

<style scoped>
body {
	background-color: #fdfbe7;
}
main {
	font-family: Arial, sans-serif;
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100vh;
	margin: 0;
	background-color: #fdfbe7;
}

.circle-chart {
	position: relative;
	width: 300px;
	height: 300px;
	border-radius: 50%;
	background: conic-gradient(
		#f0a500 0 90deg,
		#0088cc 90deg 180deg,
		#ff6347 180deg 270deg,
		#90ee90 270deg 360deg
	);
	border: 2px solid #000;
	display: flex;
	justify-content: center;
	align-items: center;
}

.circle-chart .center-circle {
	position: absolute;
	width: 120px;
	height: 120px;
	border-radius: 50%;
	background: #fdfbe7;
	border: 2px solid #000;
	display: flex;
	justify-content: center;
	align-items: center;
	text-align: center;
	font-size: 14px;
	font-weight: bold;
	line-height: 1.5;
	z-index: 1; /* 將同心圓置於最上層 */
}

.section {
	position: absolute;
	width: 50%;
	height: 50%;
	display: flex;
	justify-content: center;
	align-items: center;
	text-align: center;
	font-size: 14px;
	font-weight: bold;
}

.section.top-left {
	top: 25%;
	left: 25%;
	transform: translate(-50%, -50%);
}

.section.top-right {
	top: 25%;
	right: 25%;
	transform: translate(50%, -50%);
}

.section.bottom-left {
	bottom: 25%;
	left: 25%;
	transform: translate(-50%, 50%);
}

.section.bottom-right {
	bottom: 25%;
	right: 25%;
	transform: translate(50%, 50%);
}

.circle-lines {
	position: absolute;
	width: 100%;
	height: 100%;
}

.circle-lines::before,
.circle-lines::after {
	content: "";
	position: absolute;
	width: 2px;
	height: 100%;
	background: #000;
	top: 0;
	left: 50%;
	transform: translateX(-50%);
	z-index: 0; /* 確保分割線在同心圓之下 */
}

.circle-lines::after {
	width: 100%;
	height: 2px;
	top: 50%;
	left: 0;
	transform: translateY(-50%);
}
</style>
