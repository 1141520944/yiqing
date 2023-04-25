<template>
	<view>
		<view class="title" :class="topNavStyle.class" :style="topNavStyle.style">
			<view class="flex_col">
				<view class="box1"></view>
				<view class="flex_grow flex_col flex_center tab">
					<view v-for="(item,index) in topNavArr" :key="index"
					 :class="{ 'active':topNavIndex==index }" 
					 :data-index="index" @tap="changeTopNav">{{item}}</view>
				</view>
				<view class="box1 align_r">
					
				</view>
			</view>
		</view>
		<card-swiper></card-swiper>
		<view>
			<uni-section title="疫情核酸小程序" subTitle="其他一切因使用本小程序而引致的任何意外、疏忽、合约毁坏、隐秘汇漏、诽谤、版权或知识产权侵犯及其所造成的损失，本小程序概不负责，亦不承担任何法律责任" type="line">
			<uni-notice-bar show-icon scrollable
							text="抗击疫情,人人有责!" />
			</uni-section>
		</view>
		
	
		<view>
			<custom-tabs type="c1" :value="value" @change="changeIndex">
			    <custom-tab-pane label="小程序操作说明" name="c1_1">
					<uni-card title="小程序操作说明" extra="开发者著">
						<view class="test-component flex-center" >
						      <hm-news-card :options="options"></hm-news-card>
						</view>
					</uni-card>
				</custom-tab-pane>
			    <custom-tab-pane label="登录" name="c1_2">
					<view class="flex-center" style="padding-top: 20rpx;">
					  <zkml-button bgColor="#5866fa" title="特殊用户登录" titleColor='#ffffff' @btnClick="click"></zkml-button>
					</view>
				</custom-tab-pane>
			</custom-tabs>
		</view>
		
		
	</view>
</template>

<script>
	import cardSwiper from "@/components/helang-cardSwiper/helang-cardSwiper"
	import HmNewsCard from '@/components/hm-news-card/index.vue'
	import zkmlButton from '@/components/zkml-Button/zkml-Button.vue'
	export default {
		components:{
			cardSwiper,
			HmNewsCard,
			zkmlButton
		},
		data() {
			return {
				value:0,
				topNavIndex:0,
				topNavArr:['疫情小程序'],
				pageScrollTop:0,	// 页面滚动距离
				options: {
				            img: '../../static/docker.png',
				            title: '说明',
				            author: '郝同学',
				            date: '2022.11.24',
				            summary
							:`（1）点击个人主页填写信息，每个人的微信仅仅支持一个账号信息,请认真对待，不要占用别人的学号。如果填写提交信息出现学号被占用请联系管理员老师后台手动解决（2）操作员用户和老师请登录对应的账户进行操作 （3）感谢您的使用,如果您也想在了解这个小程序的源码请联系我`,
				            comments: "14 评论",
				            likes: "254 喜欢",
				            url: '',
				            showComments: true,
				            showLikes: true
				          }
			}
		},
		computed:{
			onClick: function(e) {
			        console.log('onClick');
			      },
			topNavStyle(){
				let r = this.pageScrollTop  / 100;
				return {
					"class":r>=0.85?'style2':'',
					"style":`background-color: rgba(80, 100, 235,${r>=1?1:r});`
				}
			}
		},
		onLoad() {
			
		},
		// 页面滚动监听
		onPageScroll(e){
			this.pageScrollTop = Math.floor(e.scrollTop);
		},
		methods: {
			changeIndex:function(){
				
			},
			click:function(){
				uni.navigateTo({
					url:"/pages/login/login"
				})
			},
			// 顶部导航改变 
			changeTopNav(e){
				this.topNavIndex = e.currentTarget.dataset.index
			},
			// 去搜索
			toSearch(){
				// console.log("tosearch")
			}
		}
	}
</script>

<style lang="scss">
	@import "lib/global.scss";
	.flex-center{
		display: flex;
		flex-direction: row;
		justify-content: center;
		
	}
	
	page{
		background-color: #fff;
	}
	/* 标题栏 */
	.title{
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: auto;
		padding-top: var(--status-bar-height);
		z-index: 10;
		background-color: rgba(66,185,131,0);
		color: rgba(255,255,255,0.8);
		
		&>view{
			height: 44px;
		}
		
		.box1{
			width: 60rpx;
			margin: 0 40rpx;
			font-size: 36rpx;
		}
		
		
		.tab{
			&>view{
				margin: 0 30rpx;
				line-height: 64rpx;
				font-size: 36rpx;
				position: relative;
				letter-spacing: 0;
				transition: transform 0.3s ease-in-out 0s;
				transform: scale(1,1);
				
				&.active{
					// color: rgba(255,255,255,1);
					transform: scale(1.15,1.15);
				}
			}
		}
		
		&.style2{
			color: #FFF;
			background-color:rgba(80,100,235,255);
			
			.tab{
				&>view{
					&.active{
						color: #FFF;
					}
				}
			}
		}
	}

</style>