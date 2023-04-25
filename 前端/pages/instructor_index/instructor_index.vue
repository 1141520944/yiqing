<!-- <template>
	<view>
		<h1>instructor_index</h1>
		<view>{{token}}--{{uid}}--{{name}}</view>
		<button v-on:click="test()">测试</button>
	</view>
</template> -->
<template>
	<view>
		<view>
			<uni-section class="mb-10" v-bind:title="name" sub-title="辅导员,您好" type="line"></uni-section>
		</view>
		<view class="flex-center" style="padding-top: 10rpx;">
		  <zkml-button bgColor="#5866fa" title="扫一扫" titleColor='#ffffff' @btnClick="scan()"></zkml-button>
		</view>
		
		<view>
			<custom-tabs type="c1" :value="value" @change="changeIndex">
			    <custom-tab-pane label="二维码的同学个人信息" name="c1_1">
					<view style="padding-top: 10rpx;">
						<uni-card title="信息" v-bind:extra="student.student_name" v-if="show ==0">
							<text>没有接受到信息</text>
						</uni-card>
						<uni-card title="信息" v-bind:extra="student.student_name" v-if="show ==1">
							<view><text>班级:{{student.class_name}}</text></view>
							<view><text>班级总人数:{{result.class_number}}</text></view>
							<view><text>班级已做核酸人数:{{result.class_rnumber}}</text></view>
							<view><text>该学生的辅导员下学生数目：{{result.instructor_number}}</text></view>
							<view><text>该学生的辅导员下已做核酸学生数目：{{result.instructor_rnumber}}</text></view>
							<view><text>学院:{{student.college}}</text></view>
							<view><text>年级:{{student.class_grade}}</text></view>
							<view><text>姓名:{{student.student_name}}</text></view>
							<view><text>学号:{{student.student_number}}</text></view>
							<view v-if="result.student_rnumber==1"><text>该学生已做核酸</text></view>
							<view v-if="result.student_rnumber==0"><text>该学生未做核酸</text></view>
						</uni-card>
					</view>
				</custom-tab-pane>
			    <custom-tab-pane label="操作说明" name="c1_2">
					<uni-card title="操作说明" extra="开发者">
						<view><text>1.正常登录-当登录失效时，请重新登录</text></view>
						<view><text>2.点击</text></view>
						<view><text>QQ:1141520944</text></view>
						<view><text>wx：hjz1141520</text></view>
						<view><text>添加请备注原因</text></view>
						<view><text>您的使用是对我最大的支持</text></view>
					</uni-card>
				</custom-tab-pane>
			    <custom-tab-pane label="常见问题" name="c1_3">
						<uni-card title="常见问题" extra="开发者">
							<view><text>常见问题:</text></view>
							<view><text>1.正常登录一段时间后，当认证令牌失效，使得无法正常添加记录，会有提示登录失效,解决办法--请您返回重新登陆</text></view>
							<view><text>2.为了放置您随意多次请求导致冗余大量的数据，我们设置按钮只能按一次，当按钮变成红色之后您需要重写扫描二维码才能进入按钮界面</text></view>
							<view><text>3.当地址您不选择时默认为 '其他'</text></view>
							<view><text>4.如果发现其他问题请联系我，感谢您的使用</text></view>
						</uni-card>
				</custom-tab-pane>
			</custom-tabs>
		</view>
	</view>
</template>

<script>
	import zkmlButton from '@/components/zkml-Button/zkml-Button.vue'
	export default {
		components: {
		    zkmlButton
		    },
		data() {
			return {
				value:0,
				token:"",
				uid:"",
				name:"",
				show:"",
				student:{
					id:"",
					class_id:"",
					college:"",
					create_time:"",
					instructor_id:"",
					open_id:"",
					phone:"",
					student_name:"",
					student_number:"",
					todatTime:"",
					update_time:"",
					user_id:"",
					class_name:"",
					class_grade:"",
					
					
				},
				post_data:{
					class_id: "1529377326108672",
					college: "软件学院",
					stdudent_number: "20201611312",
					student_name: "郝泾钊",
					u_id: "1531034491424768"
				},
				result:{
					class_name: "",
					class_number: "",
					class_rnumber: "",
					instructor_number: "",
					instructor_rnumber: "",
					student_college: "",
					student_gard: "",
					student_id_number: "",
					student_name: "",
					student_number: "",
					student_rnumber: ""
				}
			}
		},
		methods: {
			changeIndex:function(){
				// this.value=e.value
			},
			formatDate: function(value) {
			              var index=value.lastIndexOf("T");
			                value=value.substring(0,index);
			                 // console.log(value);
			                return value;
			          },
			getInfo:function(){
				this.$myHttp({
				    url: '/student/getOne?uid='+this.student.id,
					method:"get"
				}).then((res) => {
				var result = res.data
					var result = res.data
					// console.log(result)
					let that = this
					this.show=1
					this.student.class_id=result.data.class_id
					this.student.college=result.data.college
					this.student.create_time=result.data.create_time
					this.student.instructor_id=result.data.instructor_id
					this.student.open_id=result.data.open_id
					this.student.phone=result.data.phone
					this.student.student_name=result.data.student_name
					this.student.student_number=result.data.student_number
					this.student.todatTime= this.formatDate(result.data.todatTime)
					this.student.today_state=result.data.today_state
					this.student.update_time=result.data.update_time
					this.student.user_id=result.data.user_id
					// console.log(this.student)
					//查看班级
					this.$myHttp({
					    url: '/student/class/one?class_id='+this.student.class_id,
						method:"get",
					}).then((res) => {
						// console.log(res)
						var result = res.data
						this.student.class_name =result.data.name
						this.student.class_grade=result.data.grade
					})
					this.select()
				})
			},
			scan:function(){
				uni.scanCode({
					success: (res) => {
						this.student.id = res.result
						this.getInfo()
						// this.select()
					}
				})
			},
			select:function(){
					    this.$myHttp({
					        url: '/instructor/nucleicAcid/number',
							method:"post",
							data:{
								"class_id": this.student.class_id,
								"college": this.student.college,
								"stdudent_number": this.student.student_number,
								"student_name": this.student.student_name,
								"u_id": this.student.id
							},
							token:this.token
					    }).then((res) => {
							var result = res.data
							// console.log(res)
							this.result.class_name=result.data.class_name;
							this.result.class_number=result.data.class_number;
							this.result.class_rnumber= result.data.class_rnumber;
							this.result.instructor_number= result.data.instructor_number;
							this.result.instructor_rnumber= result.data.instructor_rnumber;
							this.result.student_college= result.data.student_college;
							this.result.student_gard= result.data.student_gard;
							this.result.student_id_number= result.data.student_id_number;
							this.result.student_name= result.data.student_name;
							this.result.student_number=result.data.student_number;
							this.result.student_rnumber= result.data.student_rnumber;
							// console.log(this.result)
							// console.log(this.result)
					    })
			},
		},
		created() {
			this.token= uni.getStorageSync('token');
			this.uid= uni.getStorageSync('uid');
			this.name= uni.getStorageSync('user_name');
		}
	}
</script>

<style>
	.flex-center{
		display: flex;
		flex-direction: row;
		justify-content: center;
		
	}
</style>
