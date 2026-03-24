<script lang="js">

export default {

  name: "views-upload-view",


  data() {
    const fileList = []
    return { fileList }
  },


  methods: {

    handlePreview() { },

    handleRemove() { },

    beforeRemove() { },

    handleExceed() { },

  }

}

</script>

<style></style>

<template>
  <div class="about">
    <h1> Todo: upload </h1>



    <el-upload v-model:file-list="fileList" class="upload-demo" action="/files/application/img-type-2/0000/unnamed.gif"
      multiple :on-preview="handlePreview" :on-remove="handleRemove" :before-remove="beforeRemove" :limit="3"
      :on-exceed="handleExceed">
      <el-button type="primary">Click to upload</el-button>
      <template #tip>
        <div class="el-upload__tip">
          jpg/png files with a size less than 500KB.
        </div>
      </template>
    </el-upload>


  </div>
</template>
