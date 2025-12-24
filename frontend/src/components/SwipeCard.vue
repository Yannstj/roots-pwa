<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  question: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['swipe'])

const startX = ref(0)
const currentX = ref(0)
const isDragging = ref(false)

const translateX = computed(() => currentX.value - startX.value)
const rotation = computed(() => translateX.value * 0.1)

const cardStyle = computed(() => ({
  transform: `translateX(${translateX.value}px) rotate(${rotation.value}deg)`,
  transition: isDragging.value ? 'none' : 'transform 0.3s ease-out',
}))

const swipeDirection = computed(() => {
  if (translateX.value > 50) return 'right'
  if (translateX.value < -50) return 'left'
  return null
})

function handleStart(event) {
  isDragging.value = true
  const clientX = event.touches ? event.touches[0].clientX : event.clientX
  startX.value = clientX
  currentX.value = clientX
}

function handleMove(event) {
  if (!isDragging.value) return
  const clientX = event.touches ? event.touches[0].clientX : event.clientX
  currentX.value = clientX
}

function handleEnd() {
  isDragging.value = false

  if (Math.abs(translateX.value) > 100) {
    const direction = translateX.value > 0 ? 'right' : 'left'
    emit('swipe', direction)
  }

  startX.value = 0
  currentX.value = 0
}
</script>

<template>
  <div
    class="card bg-base-100 shadow-2xl w-full max-w-sm cursor-grab active:cursor-grabbing select-none overflow-hidden"
    :style="cardStyle"
    @mousedown="handleStart"
    @mousemove="handleMove"
    @mouseup="handleEnd"
    @mouseleave="handleEnd"
    @touchstart="handleStart"
    @touchmove="handleMove"
    @touchend="handleEnd"
  >
    <!-- Indicateurs de swipe -->
    <div
      v-if="swipeDirection === 'left'"
      class="absolute top-4 left-4 badge badge-error badge-lg z-10"
    >
      NON
    </div>
    <div
      v-if="swipeDirection === 'right'"
      class="absolute top-4 right-4 badge badge-success badge-lg z-10"
    >
      OUI
    </div>

    <!-- Image -->
    <figure class="aspect-[3/4] overflow-hidden">
      <img
        :src="question.image_url"
        :alt="question.question"
        class="w-full h-full object-cover pointer-events-none"
        draggable="false"
      />
    </figure>

    <!-- Question -->
    <div class="card-body p-4 shrink-0">
      <h2 class="card-title text-center text-lg justify-center">{{ question.question }}</h2>
      <p v-if="question.category" class="badge badge-outline mx-auto">{{ question.category }}</p>
    </div>
  </div>
</template>
