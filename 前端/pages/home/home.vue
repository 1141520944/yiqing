<template>
	<view>
		<view>
		     <drag-button
		      :isDock="true"
		      :existTabBar="true"
			  :dataText="button_name"
		      @btnClick="btnClick"
							></drag-button>
		 </view>
		<view class='UCenter-bg'>
			<view class='padding-top-sm' style="padding-top: 20upx">
				<image src="../../static/docker.png"></image>
			</view>
			<view class='text-xl' style="font-size: 36upx;">你好
			</view>
			<image src='https://api.szyashu.com/static/wave.gif' mode='scaleToFill' class='gif-wave'></image>
		</view>
		<view>
			<custom-tabs type="c1" :value="value" @change="changeIndex">
			    <custom-tab-pane label="个人信息" name="c1_1">
					<view style="padding-top: 10rpx;">
						<uni-card title="个人信息" v-bind:extra="student.student_name" v-if="show ==0">
							<text>您或许没有提交信息或则没有网络</text>
						</uni-card>
						<uni-card title="个人信息" v-bind:extra="student.student_name" v-if="show ==1">
							<view><text>学院：{{student.college}}</text></view>
							<view><text>年级：{{student.class_grade}}</text></view>
							<view><text>辅导员：{{student.instructor_name}}</text></view>
							<view><text>姓名：{{student.student_name}}</text></view>
							<view><text>班级：{{student.class_name}}</text></view>
							<view><text>学号：{{student.student_number}}</text></view>
							<view><text>电话：{{student.phone}}</text></view>
							<view><text>最近核酸时间：{{student.todatTime}}</text></view>
						</uni-card>
					</view>
				</custom-tab-pane>
			    <custom-tab-pane label="操作" name="c1_2">
					<view class="flex-center">
					  <zkml-button bgColor="#5866fa" title="添加/更新信息" titleColor='#ffffff' @btnClick="click"></zkml-button>
					</view>
					<!-- <view class="flex-center" style="padding-top: 10rpx;">
					  <zkml-button bgColor="#5866fa" title="查看核酸" titleColor='#ffffff' @btnClick=""></zkml-button>
					</view> -->
				</custom-tab-pane>
				<custom-tab-pane label="二维码" name="c1_3">
					<uni-card title="二维码" v-bind:extra="student.student_name" v-if="show ==1">
						<!-- 二维码 -->
						<view class="qr-box">
							<canvas canvas-id="qrcode" v-show="qrShow" style="width: 400rpx;margin: 0 auto;"/>
						</view>	
					</uni-card>
					<uni-card title="二维码" v-bind:extra="student.student_name" v-if="show !=1">
						<!-- 二维码 -->
						<view class="qr-box">
							<text>您需要刷新出刷新出个人信息才能有二维码</text>
						</view>	
					</uni-card>
				</custom-tab-pane>
			    <custom-tab-pane label="联系我们" name="c1_3">
					<uni-card title="问题反馈" extra="开发者">
						<view><text>如果您使用途中发现bug请联系郝同学，</text></view>
						<view><text>请不要情绪激动，能解决我会尽力解决</text></view>
						<view><text>QQ:1141520944</text></view>
						<view><text>wx：hjz1141520</text></view>
						<view><text>添加请备注原因</text></view>
						<view><text>您的使用是对我最大的支持</text></view>
					</uni-card>
				</custom-tab-pane>
			</custom-tabs>
		</view>
		
	</view>
</template>

<script>
	import zkmlButton from '@/components/zkml-Button/zkml-Button.vue'
	import dragButton from "@/components/drag-button/drag-button.vue";
	import uQRCode from '@/util/u-qrcode.js' //引入uqrcode.js
	export default {
		components: {
		    zkmlButton,
			dragButton
		    },
		data() {
			return {
				button_name:"刷新",
				open_id:"",
				qrShow: false,
				text1:'',
				
				value: 0,
				show:0,
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
			}
		},
		methods: {
			// 获得open_id
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
			btnClick:function(){
				if (this.open_id!=''){
					// console.log(this.open_id)
					this.getInfoPeople()
				}else{
					this.getOpenid()
				}
			},
			// 获得用户的详细信息
			getInfoPeople:function(){
				this.$myHttp({
				    url: '/student/one',
					method:"post",
					data:{
						"open_id":this.open_id,
					}
				}).then((res) => {
					var result = res.data
					// console.log(result)
					if (result.code==1009){
						//请先提交信息再刷新
						uni.showToast({
								title:"请先提交信息再刷新",
								duration: 1500,
								icon:'none'
							});
							return
					}
					var result = res.data
					// console.log(result)
					this.show=1
					this.student.class_id=result.data.student.class_id
					this.student.college=result.data.student.college
					this.student.create_time=result.data.student.create_time
					this.student.instructor_id=result.data.student.instructor_id
					this.student.open_id=result.data.student.open_id
					this.student.phone=result.data.student.phone
					this.student.role_id=result.data.student.role_id
					this.student.student_name=result.data.student.student_name
					this.student.student_number=result.data.student.student_number
					this.student.todatTime=this.formatDate(result.data.student.todatTime)
					this.student.today_state=result.data.student.today_state
					this.student.update_time=result.data.student.update_time
					this.student.user_id=result.data.student.user_id
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
			//**生成二维码**//
			qrFun: function(text) {
				this.qrShow = true
				uQRCode.make({
					canvasId: 'qrcode',
					componentInstance: this,
					text: text,
					size: 150,
					margin: 0,
					backgroundColor: '#ffffff',
					foregroundColor: '#000000',
					fileType: 'jpg',
					errorCorrectLevel: uQRCode.errorCorrectLevel.H,
					success: res => {}
				})
			},
			changeIndex(e) {
			    // console.log('选中:', e)
				if (e.value==2){
					if (this.student.user_id!=''){
						this.qrFun(this.student.user_id)
					}
				}
			},
				formatDate: function(value) {
							  var index=value.lastIndexOf("T");
								value=value.substring(0,index);
								 // console.log(value);
								return value;
						  },
				click:function(){
					uni.navigateTo({
						url:'/pages/peopleInfo/peopleInfo'
					})
				}
		},
	created() {
		if (this.open_id==''){
			this.getOpenid()
		}else{
			return
		}
	},
	}
</script>

<style>
	
	.qr-box {
		width: 400rpx;
		height: 460rpx;
		margin: 0 auto;
		margin-top: 20rpx;
	}
	.flex-center{
		display: flex;
		flex-direction: row;
		justify-content: center;
		
	}
	
	.UCenter-bg {
	  /* background-image: url('https://api.szyashu.com/static/index.jpg'); */
	  /* background-size: cover; */
	  background-color: #66ccff;
	  height: 400rpx;
	  display: flex;
	  justify-content: center;
	  padding-top: 40rpx;
	  overflow: hidden;
	  position: relative;
	  flex-direction: column;
	  align-items: center;
	  color: #fff;
	  font-weight: 300;
	  text-shadow: 0 0 3px rgba(0, 0, 0, 0.3);
	}
	
	.UCenter-bg text {
	  opacity: 0.8;
	}
	
	.UCenter-bg image {
	  width: 200rpx;
	  height: 200rpx;
	}
	
	.UCenter-bg .gif-wave{
	  position: absolute;
	  width: 100%;
	  bottom: 0;
	  left: 0;
	  z-index: 99;
	  mix-blend-mode: screen;  
	  height: 100rpx;   
	}
</style>