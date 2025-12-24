// Utilise l'IP/hostname actuel pour que ça marche sur mobile
const API_BASE_URL = `http://${window.location.hostname}:8080/api`

export async function getQuestion() {
  const response = await fetch(`${API_BASE_URL}/question`)
  if (!response.ok) {
    throw new Error('Erreur lors de la récupération de la question')
  }
  return response.json()
}

export async function submitSwipe(questionId, direction) {
  const response = await fetch(`${API_BASE_URL}/swipe`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      question_id: questionId,
      direction: direction,
    }),
  })
  if (!response.ok) {
    throw new Error('Erreur lors de la soumission du swipe')
  }
  return response.json()
}

export async function healthCheck() {
  const response = await fetch(`http://${window.location.hostname}:8080/health`)
  return response.json()
}
