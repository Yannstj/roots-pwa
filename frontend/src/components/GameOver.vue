<script setup>
defineProps({
  score: {
    type: Number,
    default: 0,
  },
})

const emit = defineEmits(['restart'])

function handleRestart() {
  emit('restart')
}
</script>

<template>
  <div class="flex flex-col items-center justify-center h-full gap-8 p-4">
    <div class="text-center">
      <h1 class="text-4xl font-bold text-error mb-4">Game Over</h1>
      <p class="text-2xl">Score: {{ score }}</p>
    </div>

    <button class="btn btn-primary btn-lg" @click="handleRestart">Rejouer</button>
  </div>
</template>
