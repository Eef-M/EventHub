<template>
  <div>
    <div v-if="paymentStore.clientSecret && stripe">
      <div id="payment-element">
      </div>
      <button id="submit-button" @click="handleSubmit" :disabled="paymentStore.loading || submitting"
        :class="{ 'loading': submitting }">
        {{ submitting ? 'Processing...' : paymentStore.loading ? 'Loading...' : 'Pay Now' }}
      </button>
    </div>
    <div v-else-if="paymentStore.loading">
      <div class="loading-state">
        Loading payment form...
      </div>
    </div>
    <div v-else-if="error">
      <div class="error-state">
        <p class="error">{{ error }}</p>
        <button @click="retryPayment" class="retry-button">Retry</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick, watch } from 'vue'
import { loadStripe } from '@stripe/stripe-js'
import type { Stripe, StripeElements } from '@stripe/stripe-js'
import { usePaymentStore } from '@/stores/paymentStore'

interface Props {
  eventId?: string
  amount?: number
  currency?: string
  returnUrl?: string
}

const props = withDefaults(defineProps<Props>(), {
  eventId: 'REPLACE_WITH_EVENT_UUID',
  amount: 5000,
  currency: 'usd',
  returnUrl: '/payment-success'
})

const paymentStore = usePaymentStore()
const submitting = ref(false)
const error = ref<string | null>(null)
const stripe = ref<Stripe | null>(null)
const elements = ref<StripeElements | null>(null)
const paymentElement = ref<any>(null)

onMounted(async () => {
  try {
    const stripePublishableKey = import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY
    if (!stripePublishableKey) {
      throw new Error('Stripe publishable key is not configured')
    }

    stripe.value = await loadStripe(stripePublishableKey)
    if (!stripe.value) {
      throw new Error('Failed to load Stripe')
    }

    await initializePayment()
  } catch (err) {
    console.error('Stripe initialization error:', err)
    error.value = err instanceof Error ? err.message : 'Failed to initialize payment system'
  }
})

watch(() => paymentStore.clientSecret, async (newClientSecret) => {
  if (newClientSecret && stripe.value) {
    await initializeStripeElements()
  }
})

const initializePayment = async () => {
  try {
    error.value = null
    await paymentStore.createPayment({
      event_id: props.eventId,
      amount: props.amount,
      currency: props.currency
    })
  } catch (err) {
    console.error('Payment initialization error:', err)
    error.value = 'Failed to initialize payment'
  }
}

const initializeStripeElements = async () => {
  if (!stripe.value || !paymentStore.clientSecret) return

  try {
    await nextTick()
    if (paymentElement.value) {
      paymentElement.value.destroy()
    }

    const appearance = {
      theme: 'stripe' as const,
      variables: {
        colorPrimary: '#635bff',
        colorBackground: '#ffffff',
        colorText: '#30313d',
        colorDanger: '#df1b41',
        fontFamily: 'system-ui, sans-serif',
        spacingUnit: '4px',
        borderRadius: '4px',
      }
    }

    elements.value = stripe.value.elements({
      clientSecret: paymentStore.clientSecret,
      appearance
    })

    paymentElement.value = elements.value.create('payment', {
      layout: 'tabs'
    })

    paymentElement.value.mount('#payment-element')

  } catch (err) {
    console.error('Stripe Elements initialization error:', err)
    error.value = 'Failed to initialize payment form'
  }
}

const handleSubmit = async () => {
  if (!stripe.value || !elements.value) {
    error.value = 'Payment system not ready'
    return
  }

  submitting.value = true
  error.value = null

  try {
    const { error: stripeError } = await stripe.value.confirmPayment({
      elements: elements.value,
      confirmParams: {
        return_url: `${window.location.origin}${props.returnUrl}`
      }
    })

    if (stripeError) {
      console.error('Payment confirmation error:', stripeError)
      error.value = stripeError.message || 'Payment failed'
    }
  } catch (err) {
    console.error('Payment submission error:', err)
    error.value = 'An unexpected error occurred during payment'
  } finally {
    submitting.value = false
  }
}

const retryPayment = () => {
  error.value = null
  submitting.value = false

  if (paymentElement.value) {
    paymentElement.value.destroy()
    paymentElement.value = null
  }

  initializePayment()
}
</script>

<style scoped>
#payment-element {
  margin-bottom: 20px;
  min-height: 60px;
}

button {
  width: 100%;
  padding: 14px 20px;
  background-color: #635bff;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  letter-spacing: 0.025em;
}

button:hover:not(:disabled) {
  background-color: #5a52e8;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(99, 91, 255, 0.4);
}

button:disabled {
  background-color: #a8a8a8;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

button.loading {
  background-color: #5a52e8;
  position: relative;
}

.loading-state {
  text-align: center;
  padding: 40px 20px;
  color: #666;
  font-size: 16px;
}

.error-state {
  text-align: center;
  padding: 20px;
}

.error {
  color: #df1b41;
  margin: 0 0 16px 0;
  padding: 12px 16px;
  background-color: #fef7f7;
  border: 1px solid #fecaca;
  border-radius: 6px;
  font-size: 14px;
}

.retry-button {
  background-color: #6b7280;
  padding: 10px 20px;
  font-size: 14px;
  font-weight: 500;
}

.retry-button:hover {
  background-color: #4b5563;
}

/* Loading animation */
@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

button.loading::after {
  content: '';
  position: absolute;
  width: 16px;
  height: 16px;
  margin: auto;
  border: 2px solid transparent;
  border-top-color: #ffffff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
}
</style>