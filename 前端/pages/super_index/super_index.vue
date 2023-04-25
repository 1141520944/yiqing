<template>
	<view>
		<view>
			<uni-section class="mb-10" v-bind:title="name" sub-title="操作员,您好" type="line"></uni-section>
		</view>
	<!-- 	<button v-on:click="scan()">扫一扫</button> -->
				<!-- <view>{{student.id}}</view> -->
		<view class="flex-center" style="padding-top: 10rpx;">
		  <zkml-button bgColor="#5866fa" title="扫一扫" titleColor='#ffffff' @btnClick="scan()"></zkml-button>
		</view>
		
		<view>
			<custom-tabs type="c1" :value="value" @change="changeIndex">
			    <custom-tab-pane label="二维码的同学个人信息" name="c1_1">
					<view style="padding-top: 10rpx;">
						<uni-card title="信息" v-bind:extra="up_data" v-if="show ==0">
							<text>没有接受到信息</text>
						</uni-card>
						<uni-card title="信息" v-bind:extra="up_data" v-if="show ==1">
							<view><text>学院：{{student.college}}</text></view>
							<view><text>年级：{{student.class_grade}}</text></view>
							<view><text>辅导员：{{student.instructor_name}}</text></view>
							<view><text>姓名：{{student.student_name}}</text></view>
							<view><text>班级：{{student.class_name}}</text></view>
							<view><text>学号：{{student.student_number}}</text></view>
							<view><text>电话：{{student.phone}}</text></view>
							<view v-if="!(oldtime=='')"><text>上次核酸时间：{{oldtime}}</text></view>
							<view><text>记录最近核酸时间：{{student.todatTime}}</text></view>
							<view class="ui-all">
								<view class="ui-list right">
									<text>位置</text>
									<picker @change="bindAdressChange" mode='selector' range-key="name" :value="value" :range="address_data">
										<view class="picker">
											{{address_data[index_address].name}}
										</view>
									</picker>
								</view>
							</view>
							<view class="flex-center" style="padding-top: 10rpx;">
							  <zkml-button v-bind:bgColor="buttom_color" v-bind:title="button_name" titleColor='#ffffff' @btnClick="add()" ></zkml-button>
							</view>
						</uni-card>
					</view>
				</custom-tab-pane>
			    <custom-tab-pane label="操作说明" name="c1_2">
					<uni-card title="操作说明" extra="开发者">
						<view><text>1.正常登录-当登录失效时，请重新登录</text></view>
						<view><text>2.点击扫一扫，扫描学生的二维码，选择地点位置</text></view>
						<view><text>3.当数据修改成功，左上角会由‘未提交’变为‘提交成功’ 并且会出现该学生最近记录的两次核酸时间，即当次核酸时间和上次核酸时间</text></view>
						<view><text>4.再次扫描新的学生二维码，重复以上操作</text></view>
						<view><text>5.设置最短登录间隔时间差为2小时,按照提示操作，感谢您的使用</text></view>
					</uni-card>
				</custom-tab-pane>
			    <custom-tab-pane label="常见问题" name="c1_3">
						<uni-card title="常见问题" extra="开发者">
							<view><text>常见问题:</text></view>
							<view><text>1.正常登录一段时间后，当认证令牌失效，使得无法正常添加记录，会有提示登录失效,解决办法--请您返回重新登陆</text></view>
							<view><text>2.为了放置您随意多次请求导致冗余大量的数据，我们设置按钮只能按一次，当按钮变成红色之后您需要重写扫描二维码才能进入按钮界面</text></view>
							<view><text>3.当地址您不选择时默认为 上传的值为'其他' 选项</text></view>
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
				token:"",
				uid:"",
				name:"",

				uid:"",
				number:"",
				value:0,
				show:0,
				oldtime:"",
				up_data:"未提交",
				button_name:"添加核酸记录",
				button_go:true,
				buttom_color:"#5866fa",
				student:{
					class_id: "",
					college: "",
					create_time: "",
					instructor_id: "",
					open_id: "",
					phone: "",
					role_id: 0,
					student_name: "无",
					student_number: "",
					todatTime: "",
					today_state: 0,
					update_time: "",
					user_id: "",
					class_grade: "",
					class_name: "",
					class_number: "",
					class_state_number: "",
					instructor_name:""
				},
				
				index_address:0,
				address_data: [{
					id: 1,
					name: '未填写 ',
				}, {
					id: 2,
					name: '大操场',
				},
				{
					id: 3,
					name: '小操场',
				},
				{
					id: 4,
					name: '其他'
				}
				],
			}
		},
		methods: {
			bindAdressChange:function(e){
				this.index_address=e.detail.value
			},
			changeIndex:function(e){
				
			},
			formatDate: function(value) {
			              var index=value.lastIndexOf("T");
			                value=value.substring(0,index);
			                 // console.log(value);
			                return value;
			          },
			//得到用户的信息
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
					this.student.role_id=result.data.role_id
					this.student.student_name=result.data.student_name
					this.student.student_number=result.data.student_number
					this.student.todatTime= this.formatDate(result.data.todatTime)
					this.student.today_state=result.data.today_state
					this.student.update_time=result.data.update_time
					this.student.user_id=result.data.user_id
					
					//查看班级
					this.$myHttp({
					    url: '/student/class/one?class_id='+this.student.class_id,
						method:"get",
					}).then((res) => {
						// console.log(res)
						var result = res.data
						this.student.class_name =result.data.name
						this.student.class_grade=result.data.grade
						this.student.class_number=result.data.number
						this.student.class_state_number=result.data.state_number
					})
					//查看辅导员
					//查看班级
					this.$myHttp({
					    url: '/student/instructor/one?instructor_id='+this.student.instructor_id,
						method:"get",
					}).then((res) => {
						// console.log(res)
						var result = res.data
						this.student.instructor_name =result.data.name
						
					})

					})
			},
			//扫描二位码
			scan() {
				uni.scanCode({
					success: (res) => {
						// console.log(res);
						this.up_data="未提交"
						this.buttom_color="#5866fa"
						this.student.id = res.result
						this.getInfo()
						this.button_go=true
						this.button_name="添加核酸记录"
						this.oldtime=''
					}
				})
			},
			add:function(){
				this.button_name="按钮已失效"
				this.buttom_color="red"
				if(!this.button_go){
					return
				}
				this.button_go=false
				if (this.index_address==0){
					this.index_address=3
				}
				this.oldtime=this.student.todatTime
				this.$myHttp({
				    url: '/super/nucleicAcid/add',
					method:"post",
					data:{
						"adderss": this.address_data[this.index_address].name,
						"user_id": this.student.user_id
					},
					token:this.token
				}).then((res) => {
					var result = res.data
					if (result.code==1007){
						return uni.showToast({
								title:"登录失效",
								duration: 1300,
								icon:'error'
							})
					}else if(result.code==1000) {
							this.getInfo()
							this.up_data="提交成功"
					}
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

<style lang="less">
	.flex-center{
		display: flex;
		flex-direction: row;
		justify-content: center;
		
	}
	.container {
		display: block;
	}

	.ui-all {
		padding: 20rpx 40rpx;

		.avatar {
			width: 100%;
			text-align: left;
			padding: 20rpx 0;
			border-bottom: solid 1px #f2f2f2;
			position: relative;

			.imgAvatar {
				width: 140rpx;
				height: 140rpx;
				border-radius: 50%;
				display: inline-block;
				vertical-align: middle;
				overflow: hidden;

				.iavatar {
					width: 100%;
					height: 100%;
					display: block;
				}
			}

			text {
				display: inline-block;
				vertical-align: middle;
				color: #8e8e93;
				font-size: 28rpx;
				margin-left: 40rpx;
			}

			&:after {
				content: ' ';
				width: 20rpx;
				height: 20rpx;
				border-top: solid 1px #030303;
				border-right: solid 1px #030303;
				transform: rotate(45deg);
				-ms-transform: rotate(45deg);
				/* IE 9 */
				-moz-transform: rotate(45deg);
				/* Firefox */
				-webkit-transform: rotate(45deg);
				/* Safari 和 Chrome */
				-o-transform: rotate(45deg);
				position: absolute;
				top: 85rpx;
				right: 0;
			}
		}

		.ui-list {
			width: 100%;
			text-align: left;
			padding: 20rpx 0;
			border-bottom: solid 1px #f2f2f2;
			position: relative;

			text {
				color: #4a4a4a;
				font-size: 28rpx;
				display: inline-block;
				vertical-align: middle;
				min-width: 150rpx;
			}

			input {
				color: #030303;
				font-size: 30rpx;
				display: inline-block;
				vertical-align: middle;
			}
			button{
				color: #030303;
				font-size: 30rpx;
				display: inline-block;
				vertical-align: middle;
				background: none;
				margin: 0;
				padding: 0;
				&::after{
					display: none;
				}
			}
			picker {
				width: 90%;
				color: #030303;
				font-size: 30rpx;
				display: inline-block;
				vertical-align: middle;
				position: absolute;
				top: 30rpx;
				left: 150rpx;
			}

			textarea {
				color: #030303;
				font-size: 30rpx;
				vertical-align: middle;
				height: 150rpx;
				width: 100%;
				margin-top: 50rpx;
			}

			.place {
				color: #999999;
				font-size: 28rpx;
			}
		}

		.right:after {
			content: ' ';
			width: 20rpx;
			height: 20rpx;
			border-top: solid 1px #030303;
			border-right: solid 1px #030303;
			transform: rotate(45deg);
			-ms-transform: rotate(45deg);
			/* IE 9 */
			-moz-transform: rotate(45deg);
			/* Firefox */
			-webkit-transform: rotate(45deg);
			/* Safari 和 Chrome */
			-o-transform: rotate(45deg);
			position: absolute;
			top: 40rpx;
			right: 0;
		}

		.save {
			background: #030303;
			border: none;
			color: #ffffff;
			margin-top: 40rpx;
			font-size: 28rpx;
		}
	}
</style>

