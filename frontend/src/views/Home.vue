<script setup>
import { ref, onMounted } from 'vue'
import SwipeCard from '@/components/SwipeCard.vue'
import GameOver from '@/components/GameOver.vue'
import { getQuestion, submitSwipe } from '@/services/api'

const currentQuestion = ref(null)
const score = ref(0)
const isGameOver = ref(false)
const isLoading = ref(true)
const error = ref(null)

async function loadQuestion() {
  isLoading.value = true
  error.value = null
  try {
    currentQuestion.value = await getQuestion()
  } catch (e) {
    error.value = e.message
  } finally {
    isLoading.value = false
  }
}

async function handleSwipe(direction) {
  if (!currentQuestion.value) return

  try {
    const result = await submitSwipe(currentQuestion.value.id, direction)

    if (result.correct) {
      score.value++
      await loadQuestion()
    } else {
      isGameOver.value = true
    }
  } catch (e) {
    error.value = e.message
  }
}

function handleRestart() {
  score.value = 0
  isGameOver.value = false
  loadQuestion()
}

onMounted(() => {
  loadQuestion()
})
</script>

<template>
  <div class="relative flex flex-col items-center h-full p-4 pt-16 pb-safe">
    <!-- Score -->
    <div class="absolute top-4 left-4">
      <div class="badge badge-primary badge-lg">Score: {{ score }}</div>
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="flex flex-col items-center gap-4">
      <span class="loading loading-spinner loading-lg"></span>
      <p>Chargement...</p>
    </div>

    <!-- Error -->
    <div v-else-if="error" class="alert alert-error max-w-sm">
      <span>{{ error }}</span>
      <button class="btn btn-sm" @click="loadQuestion">Réessayer</button>
    </div>

    <!-- Game Over -->
    <GameOver v-else-if="isGameOver" :score="score" @restart="handleRestart" />

    <!-- Swipe Card -->
    <SwipeCard v-else-if="currentQuestion" :question="currentQuestion" @swipe="handleSwipe" />
  </div>
</template>
