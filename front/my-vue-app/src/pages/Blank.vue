<template>
  <div></div>
</template>

<script>
export default {
    name: 'BlankPage',
    activated() {
        this.$router.replace({
            path: '/seek-taste'
        })
    },
}
</script>

<style>

</style>