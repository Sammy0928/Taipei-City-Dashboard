<template>
	<div>
		<header>
			<b>車籍資料系統</b>
			<form @submit.prevent="fetchData">
				<input
					type="text"
					v-model="licensePlate"
					placeholder="輸入車牌號碼"
				/>
				<button type="submit">搜尋</button>
			</form>
		</header>
		<main>
			<section v-if="personData">
				<div>
					<img
						:src="personData.photo_url"
						alt="車主照片"
						width="200px"
						height="200px"
					/>
					<u>{{ personData.license_plate }}</u>
				</div>
				<div class="t1">
					姓名：{{ personData.owner_name }}<br />
					手機：{{ personData.phone }}<br />
					性別：{{ personData.gender }}<br />
					出生年月日：{{ personData.birth_date }}<br />
					戶籍地址：{{ personData.address }}<br />
					車型：{{ personData.vehicle_model }}<br />
					車輛顏色：{{ personData.vehicle_color }}<br />
					車輛品牌：{{ personData.vehicle_brand }}
				</div>
			</section>
			<div class="t2" v-if="personData">
				<b>車輛資料</b><br />
				掛牌日期：{{ personData.registration_date }}<br />
				<fieldset>
					<legend>違規紀錄：</legend>
					<ul>
						<li v-for="record in violationData" :key="record.date">
							日期：{{ record.date }}，地點：{{
								record.location
							}}，描述：{{ record.description }}，罰款：{{
								record.fine_amount
							}}
						</li>
					</ul>
				</fieldset>
			</div>
		</main>
	</div>
</template>

<script>
import axios from "axios";

export default {
	data() {
		return {
			licensePlate: "", // 默認車牌號可以設定為空
			personData: null,
			violationData: [],
		};
	},
	methods: {
		async fetchData() {
			try {
				const response = await axios.get(
					"http://localhost:8080/api/v1/cardata",
					{
						params: {
							license_plate: this.licensePlate,
						},
					}
				);
				this.personData = response.data.person_data || {};
				this.violationData = response.data.violation_data || [];
			} catch (error) {
				console.error("Error fetching data:", error);
				this.personData = {};
				this.violationData = [];
			}
		},
	},
	mounted() {
		this.fetchData(); // 頁面加載時獲取默認車主信息
	},
};
</script>

<style>
.t1 {
	font-size: medium;
	line-height: 1.8;
	margin-left: 20px;
	justify-content: center;
	align-content: center;
}

.t2 {
	font-size: medium;
	line-height: 1.8;
}

fieldset {
	max-height: 200px;
	overflow-y: scroll;
	padding: 10px;
}
</style>
