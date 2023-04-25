<template>
	<view class="container">
		   <view>
		        <drag-button
		         :isDock="true"
		         :existTabBar="true"
		         @btnClick="btnClick"
					></drag-button>
		    </view>
		
		<view class="ui-all">
		<!-- 	<view class="avatar" @tap="avatarChoose">
				<view  class="imgAvatar">
					<view class="iavatar" :style="'background: url('+avater+') no-repeat center/cover #eeeeee;'"></view>
				</view>
				<text v-if="avater">修改头像</text>
				<text v-if="!avater">授权微信</text>
				<button v-if="!avater" open-type="getUserInfo" @tap="getUserInfo" class="getInfo"></button>
			</view> -->
			<view class="ui-list">
				<text>学院</text>
				<input type="text" :placeholder="value" :value="college" @input="bindCollege" placeholder-class="place" />
			</view>
			<view class="ui-list right">
				<text>辅导员</text>
				<picker @change="bindInstructorChange" mode='selector' range-key="name" :value="index" :range="instructors">
					<view class="picker">
						{{instructors[index_instructor].name}}
					</view>
				</picker>
			</view>
			<view class="ui-list right">
				<text>班级</text>
				<picker @change="bindClassChange" mode='selector' range-key="name" :value="index" :range="classes">
					<view class="picker">
						{{classes[index_class].name}}
					</view>
				</picker>
			</view>
		<!-- 	<view class="ui-list">
				<text>手机号</text>
				<input v-if="mobile" type="tel" :placeholder="value" :value="mobile" @input="bindmobile" placeholder-class="place" />
				<button v-if="!mobile" open-type="getPhoneNumber" @getphonenumber="getphonenumber" class="getInfo bun">授权手机号</button>
			</view> -->
			<view class="ui-list">
				<text>手机号</text>
				<input type="text" :placeholder="value" :value="phone" @input="bindPhone" placeholder-class="place" />
			</view>
			<view class="ui-list">
				<text>姓名</text>
				<input type="text" :placeholder="value" :value="student_name" @input="bindStudenName" placeholder-class="place" />
			</view>
			<view class="ui-list">
				<text>学号</text>
				<input type="text" :placeholder="value" :value="student_number" @input="bindStudenNumber" placeholder-class="place" />
			</view>
		<!-- 	<view class="ui-list right">
				<text>常住地</text>
				<picker @change="bindRegionChange" mode='region'>
					<view class="picker">
						{{region[0]}} {{region[1]}} {{region[2]}}
					</view>
				</picker>
			</view> -->
		<!-- 	<view class="ui-list right">
				<text>生日</text>
				<picker mode="date" :value="date" @change="bindDateChange">
					<view class="picker">
						{{date}}
					</view>
				</picker>
			</view> -->
			<!-- <view class="ui-list">
				<text>签名</text>
				<textarea :placeholder="value" placeholder-class="place" :value="description" @input="binddescription"></textarea>
			</view> -->
			<button class="save" @tap="savaInfo">保 存 修 改</button>
		</view>

	</view>
</template>

<script>
	import dragButton from "@/components/drag-button/drag-button.vue";
	export default {
	    components: {
	        dragButton
	    },
		data() {
			return {
				value: '请填写',
				open_id:"",
				classes: [{
					class_id: 1,
					name: '未填写',
				}],
				instructors:[{
					instructor_id: 1,
					name: '没有连网 ',
				}],
				index_class: 0,
				index_instructor: 0,
				region: ['请填写'],
				date: '请填写',
				avater: '',
				description: '',
				url: '',
				college: '',
				student_number:'',
				student_name:'',
				phone:'',
				mobile: '',
				headimg: '',
				result:{
					token: "",
					userID:"" ,
					userName: ""
				}

			}

		},
		methods: {// 获得open_id
			getOpenid:function(){
				uni.login({
				        provider: "weixin",
				        success: function (res) {
						uni.request({
							  method:'POST',	
							  url: 'http://106.15.65.147:8081/api/v1/wx/openid', //仅为示例，并非真实接口地址。
							  data: {
							      "app_id": "wx8d36d8370b6e82f0",
							      "code": res.code,
							      "method": "get",
							      "secret": "092d4b45d6b6c8d2b99bf82c6e23657e",
							      "url": "https://api.weixin.qq.com/sns/jscode2session" 
							  },
							  header: {
							      'content-type': 'application/json' //自定义请求头信息
							  },
							  success: (res) => {
								  var result = res.data
								  	uni.setStorageSync('open_id',result.data.open_id)
									uni.setStorageSync('session_key',result.data.session_key)
							  }
							});
				        },
				      });
					  this.open_id= uni.getStorageSync('open_id');
			},
			getClasser:function(){
				if (this.instructors[this.index_instructor].instructor_id==1){
					//请选择辅导员老师
					uni.showToast({
							title:"请选择辅导员老师",
							duration: 1500,
							icon:'none'
						});
						return
				}else{
					this.$myHttp({
						url: '/student/class/getall?instructor_id='+this.instructors[this.index_instructor].instructor_id,
						method:"get",
					}).then((res) => {
						var result = res.data
						// console.log(result)
						this.classes=result.data.classes
					})
				}
			},
			getInstructors:function(){
				// console.log('------')
				this.$myHttp({
				    url: '/student/instructor/getall',
					method:"get",
				}).then((res) => {
					var result = res.data
					// console.log(result)
					this.instructors=result.data.instructors
					
				})
			},
			btnClick:function(){
				uni.switchTab({
						url: '/pages/home/home'
					})
			},
			bindClassChange:function(e){
				this.index_class=e.detail.value
			},
			bindInstructorChange:function(e){
				this.index_instructor=e.detail.value
				// console.log(this.index_instructor)
				this.getClasser()
			},
			
			bindCollege:function(e){
				this.college=e.detail.value
			},
			bindPhone:function(e){
				this.phone=e.detail.value
			},
			bindStudenName:function(e){
				this.student_name=e.detail.value
			},
			bindStudenNumber:function(e){
				this.student_number=e.detail.value
			},
			avatarChoose() {
				let that = this;
				uni.chooseImage({
					count: 1,
					sizeType: ['original', 'compressed'],
					sourceType: ['album', 'camera'],
					success(res) {
						// tempFilePath可以作为img标签的src属性显示图片
						that.imgUpload(res.tempFilePaths);
						const tempFilePaths = res.tempFilePaths;
					}
				});
			},
			 getUserInfo () {
				  uni.getUserProfile({
			      desc: '用于完善会员资料', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
			      success: (res) => {
			       console.log(res);
				   uni.showToast({
							   title: '已授权',
							   icon: 'none',
							   duration: 2000
							   }) 
			      }
			    })
			    } ,
				 getphonenumber(e){
					if(e.detail.iv){
					  console.log(e.detail.iv) //传后台解密换取手机号
						  uni.showToast({
							   title: '已授权',
							   icon: 'none',
							   duration: 2000
							   }) 
					}
								  },
			// uni.showToast({
			// 	title: '请填写昵称',
			// 	icon: 'none',
			// 	duration: 2000
			// });	
			//提交信息
			savaInfo() {
				let that = this
				//判断是否为空
				if (this.college==""){
					uni.showToast({
						title: '请填写学院',
						icon: 'none',
						duration: 2000
					});	
					return
				}
				if (this.classes[this.index_class].class_id==1){
					uni.showToast({
						title: '请选择班级',
						icon: 'none',
						duration: 2000
					});	
					return
				}
				if (this.instructors[this.index_instructor].instructor_id==1){
					uni.showToast({
						title: '请选择辅导员',
						icon: 'none',
						duration: 2000
					});	
					return
				}
				let reg = /^[1][3,4,5,7,8,9][0-9]{9}$/
					if (!this.phone) {      //判断如果手机号为空，提示用户输入手机号
						uni.showToast({
							title: '请输入手机号',
							duration:1500,
							icon: 'none'
						})
						return
					} else if (!reg.test(this.phone)) {        //判断手机号格式时候正确
						uni.showToast({
							title: '请输入正确的手机号',
							duration:1500,
							icon: 'none'
						})
						return
					}
				if (!this.student_name){
					uni.showToast({
						title: '请输入姓名',
						duration:1500,
						icon: 'none'
					})
					return
				}
				if (!this.student_number){
					uni.showToast({
						title: '请输入学号',
						duration:1500,
						icon: 'none'
					})
					return
				}
				//请求页面
				this.$myHttp({
				    url: '/student/add',
					method:"post",
					data:{
						   "class_id": this.classes[this.index_class].class_id,
						    "college": this.college,
						    "instructor_id": this.instructors[this.index_instructor].instructor_id,
						    "open_id": this.open_id,
						    "phone": this.phone,
						    "stdudent_number": this.student_number,
						    "student_name": this.student_name
					}
				}).then((res) => {
					var result = res.data
					console.log(result)
					if (result.code==1008){
						uni.showToast({
							title: result.msg,
							duration:1500,
							icon: 'none'
						})
						return
					}
					if (result.code=1000){
						// this.result.token=result.data.token
						// this.result.userID=result.data.userID
						// this.result.userName=result.data.userName
						//加入缓存
						uni.setStorageSync('token',result.data.token)
						uni.setStorageSync('uid',result.data.userID)
						uni.setStorageSync('user_name',result.data.userName)
						uni.showToast({
							title: "提价成功",
							duration:3000,
							icon: 'success'
						})
						// 登录成功跳转到首页
						uni.switchTab({
								url: '/pages/index/index'
							})
						return
					}
					})
			},
			isPoneAvailable(poneInput) {
				var myreg = /^[1][3,4,5,7,8][0-9]{9}$/;
				if (!myreg.test(poneInput)) {
					return false;
				} else {
					return true;
				}
			},
			async updata(datas) {
				//传后台
				
			},
			imgUpload(file) {
				let that = this;
				uni.uploadFile({
					header: {
						Authorization: uni.getStorageSync('token')
					},
					url:'/api/upload/image', //需传后台图片上传接口
					filePath: file[0],
					name: 'file',
					formData: {
						type: 'user_headimg'
					},
					success: function(res) {
						var data = JSON.parse(res.data);
						data = data.data;
						that.avater = that.url + data.img;

						that.headimg = that.url + data.img;

						
					},
					fail: function(error) {
						console.log(error);
					}
				});
			},
	
		},
		created() {		
			this.getOpenid()
			//辅导员
			this.getInstructors()
		}

	}
</script>

<style lang="less">
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
