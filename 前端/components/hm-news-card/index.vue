<template>
  <div class="hm-news-card">
    <div class="container">
      <div class="box" @click="onClick">
        <img class="img" :src="options.img" />
        <span class="title">{{ options.title }}</span>
        <div class="submain">
          <span class="author">{{ options.author }}</span>
          <span class="date">{{ options.date }}</span>
        </div>
       <text class="summary">{{ options.summary }}</text>
        <!-- <div class="row_2" /> -->
  <!--      <div class="ft">
          <span class="comments" v-show="options['showComments']">{{ options.comments }}</span>
          <span class="likes" v-show="options['showLikes']">{{ options.likes }}</span>
        </div> -->
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'HmNewsCard',
  props: {
    dataId: {
      type: String,
      default: 'hm-news-card'
    },
    options: {
      type: Object,
      default: function() {
        return {
          img: '',
          title: '',
          author: '',
          date: '',
          summary:
            '',
          // comments: "14 评论",
          // likes: "254 喜欢",
          url: '',
          showComments: true,
          showLikes: true
        };
      }
    }
  },
  data() {
    return {};
  },
  methods: {
    onClick() {
      uni.navigateTo({
        url: this.options.url
      });
    }
  }
};
</script>
<style>
@import './index.response.css';
</style>
