<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
    <!-- Backdrop -->
    <div class="fixed inset-0 bg-black/50 backdrop-blur-sm transition-opacity" @click="closeModal"></div>

    <!-- Modal -->
    <div class="flex min-h-full items-center justify-center p-4">
      <div class="relative w-full max-w-2xl transform overflow-hidden rounded-2xl bg-white shadow-2xl transition-all">
        <!-- Header -->
        <div class="flex items-center justify-between border-b border-gray-200 px-6 py-4">
          <h3 class="text-2xl font-bold text-gray-900 flex items-center">
            <CreditCardIcon class="w-6 h-6 mr-3 text-purple-600" />
            Complete Payment
          </h3>
          <button @click="closeModal"
            class="rounded-lg p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 transition-colors">
            <XIcon class="w-5 h-5" />
          </button>
        </div>

        <!-- Content -->
        <div class="px-6 py-6">
          <!-- Ticket Information -->
          <div v-if="selectedTicket"
            class="mb-8 p-4 bg-gradient-to-r from-purple-50 to-blue-50 rounded-xl border border-purple-200">
            <h4 class="text-lg font-semibold text-gray-900 mb-2">{{ selectedTicket.name }}</h4>
            <p class="text-gray-600 text-sm mb-3">{{ selectedTicket.description || 'No description available' }}</p>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <!-- Quantity Selector -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Quantity</label>
                <div class="flex items-center space-x-2">
                  <button @click="decreaseQuantity" :disabled="quantity <= 1"
                    class="w-8 h-8 rounded-full bg-gray-100 hover:bg-gray-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center transition-colors">
                    <MinusIcon class="w-4 h-4" />
                  </button>
                  <span class="w-12 text-center font-semibold">{{ quantity }}</span>
                  <button @click="increaseQuantity" :disabled="quantity >= selectedTicket.quota"
                    class="w-8 h-8 rounded-full bg-gray-100 hover:bg-gray-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center transition-colors">
                    <PlusIcon class="w-4 h-4" />
                  </button>
                </div>
                <p class="text-xs text-gray-500 mt-1">Max: {{ selectedTicket.quota }}</p>
              </div>

              <!-- Price per ticket -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Price per ticket</label>
                <p class="text-2xl font-bold text-purple-600">{{ formatPrice(selectedTicket.price) }}</p>
              </div>

              <!-- Total price -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Total Amount</label>
                <p class="text-2xl font-bold text-gray-900">{{ formatPrice(totalAmount) }}</p>
              </div>
            </div>
          </div>

          <!-- Payment Form -->
          <div class="space-y-6">
            <!-- Customer Information -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label for="customerName" class="block text-sm font-medium text-gray-700 mb-2">
                  Full Name
                </label>
                <input id="customerName" v-model="customerInfo.name" type="text" required readonly
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg bg-gray-50 text-gray-600 cursor-not-allowed"
                  :placeholder="customerInfo.name || 'Enter your full name'" />
              </div>
              <div>
                <label for="customerEmail" class="block text-sm font-medium text-gray-700 mb-2">
                  Email Address
                </label>
                <input id="customerEmail" v-model="customerInfo.email" type="email" required readonly
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg bg-gray-50 text-gray-600 cursor-not-allowed"
                  :placeholder="customerInfo.email || 'Enter your email'" />
              </div>
            </div>

            <!-- Payment Method Selection -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-3">Payment Method</label>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
                <button v-for="method in paymentMethods" :key="method.id" @click="selectPaymentMethod(method.id)"
                  :class="[
                    'flex items-center space-x-3 p-4 border-2 rounded-xl transition-all duration-200',
                    selectedPaymentMethod === method.id
                      ? 'border-purple-500 bg-purple-50 ring-1 ring-purple-500'
                      : 'border-gray-200 hover:border-gray-300'
                  ]">
                  <component :is="method.icon" class="w-6 h-6 text-gray-600" />
                  <div class="text-left">
                    <p class="font-medium text-gray-900">{{ method.name }}</p>
                    <p class="text-sm text-gray-500">{{ method.description }}</p>
                  </div>
                </button>
              </div>
            </div>

            <!-- Stripe Card Element -->
            <div v-if="selectedPaymentMethod === 'card'" class="space-y-4">
              <div class="flex items-center justify-between">
                <label class="block text-sm font-medium text-gray-700">Card Information</label>
                <div v-if="cardElementMounted" class="flex items-center space-x-2">
                  <div v-if="cardComplete" class="flex items-center text-green-600 text-xs">
                    <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd"
                        d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                        clip-rule="evenodd"></path>
                    </svg>
                    Complete
                  </div>
                  <div v-else-if="!cardValid" class="text-orange-600 text-xs">
                    Enter card details
                  </div>
                </div>
              </div>

              <!-- Loading state -->
              <div v-if="!cardElementMounted"
                class="p-4 border border-gray-300 rounded-lg min-h-[50px] flex items-center justify-center">
                <div class="flex items-center space-x-2 text-gray-400 text-sm">
                  <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-gray-400"></div>
                  <span>Loading payment form...</span>
                </div>
              </div>

              <!-- Stripe card element container -->
              <div v-show="cardElementMounted" id="card-element" :class="[
                'p-4 border rounded-lg transition-all min-h-[50px]',
                cardComplete
                  ? 'border-green-300 focus-within:ring-2 focus-within:ring-green-500 focus-within:border-transparent'
                  : 'border-gray-300 focus-within:ring-2 focus-within:ring-purple-500 focus-within:border-transparent'
              ]">
              </div>

              <div id="card-errors" class="text-red-500 text-sm" role="alert"></div>
            </div>

            <!-- Other payment methods placeholder -->
            <div v-else-if="selectedPaymentMethod === 'paypal'"
              class="p-4 border-2 border-dashed border-gray-300 rounded-xl text-center text-gray-500">
              <CreditCardIcon class="w-12 h-12 mx-auto mb-2 opacity-50" />
              <p>PayPal integration coming soon</p>
            </div>

            <div v-else-if="selectedPaymentMethod === 'bank'"
              class="p-4 border-2 border-dashed border-gray-300 rounded-xl text-center text-gray-500">
              <CreditCardIcon class="w-12 h-12 mx-auto mb-2 opacity-50" />
              <p>Bank transfer integration coming soon</p>
            </div>
          </div>

          <!-- Order Summary -->
          <div class="mt-8 p-4 bg-gray-50 rounded-xl">
            <h4 class="text-lg font-semibold text-gray-900 mb-3">Order Summary</h4>
            <div class="space-y-2">
              <div class="flex justify-between">
                <span class="text-gray-600">{{ selectedTicket?.name }} × {{ quantity }}</span>
                <span class="font-medium">{{ formatPrice(totalAmount) }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-600">Processing Fee</span>
                <span class="font-medium">{{ formatPrice(processingFee) }}</span>
              </div>
              <hr class="border-gray-200">
              <div class="flex justify-between text-lg font-bold">
                <span>Total</span>
                <span class="text-purple-600">{{ formatPrice(finalTotal) }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Footer -->
        <div class="border-t border-gray-200 px-6 py-4 bg-gray-50">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-2 text-sm text-gray-500">
              <ShieldCheckIcon class="w-4 h-4" />
              <span>Secure payment powered by Stripe</span>
            </div>
            <div class="flex space-x-3">
              <Button variant="outline" @click="closeModal" :disabled="isProcessing" class="px-6 py-2 cursor-pointer">
                Cancel
              </Button>
              <Button @click="processPayment" :disabled="!canProceedPayment || isProcessing"
                class="px-8 py-2 bg-purple-600 hover:bg-purple-700 disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer">
                <div v-if="isProcessing" class="flex items-center">
                  <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                  Processing...
                </div>
                <div v-else class="flex items-center">
                  <CreditCardIcon class="w-4 h-4 mr-2" />
                  Pay {{ formatPrice(finalTotal) }}
                </div>
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { Button } from '@/components/ui/button'
import {
  CreditCardIcon,
  XIcon,
  MinusIcon,
  PlusIcon,
  ShieldCheckIcon
} from 'lucide-vue-next'
import { loadStripe } from '@/services/paymentService'
import { usePaymentStore } from '@/stores/paymentStore'
import { formatPrice, calculateTotalPrice, validateTicketPurchase } from '@/utils/ticketUtils'
import type { Ticket } from '@/types/ticket'
import type { UserInterface } from '@/types/user'

interface Props {
  isOpen: boolean
  selectedTicket: Ticket | null
  user: UserInterface | null
}

interface Emits {
  (e: 'close'): void
  (e: 'success', paymentId: string): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const paymentStore = usePaymentStore()

const quantity = ref(1)
const selectedPaymentMethod = ref('card')
const isProcessing = ref(false)
const stripe = ref<any>(null)
const cardElement = ref<any>(null)
const cardElementMounted = ref(false)
const stripeInitialized = ref(false)
const cardValid = ref(false)
const cardComplete = ref(false)

const customerInfo = computed(() => {
  if (props.user) {
    return {
      name: `${props.user.first_name} ${props.user.last_name}`.trim(),
      email: props.user.email
    }
  }
  return {
    name: '',
    email: ''
  }
})

const paymentMethods = [
  {
    id: 'card',
    name: 'Credit Card',
    description: 'Visa, Mastercard, American Express',
    icon: CreditCardIcon
  },
  {
    id: 'paypal',
    name: 'PayPal',
    description: 'Pay with your PayPal account',
    icon: CreditCardIcon
  },
  {
    id: 'bank',
    name: 'Bank Transfer',
    description: 'Direct bank transfer',
    icon: CreditCardIcon
  }
]

const priceCalculation = computed(() => {
  if (!props.selectedTicket) return { subtotal: 0, processingFee: 0, total: 0 }
  return calculateTotalPrice(props.selectedTicket.price, quantity.value)
})

const totalAmount = computed(() => priceCalculation.value.subtotal)
const processingFee = computed(() => priceCalculation.value.processingFee)
const finalTotal = computed(() => priceCalculation.value.total)

const canProceedPayment = computed(() => {
  const basicValidation = (
    customerInfo.value.name.trim() &&
    customerInfo.value.email.trim() &&
    selectedPaymentMethod.value &&
    quantity.value > 0 &&
    props.selectedTicket &&
    props.user
  )

  // Additional validation for card payments
  if (selectedPaymentMethod.value === 'card') {
    return basicValidation && cardElementMounted.value && cardElement.value && cardValid.value && cardComplete.value
  }

  return basicValidation
})

function closeModal() {
  cleanupStripeElements()
  emit('close')
  resetForm()
}

function resetForm() {
  quantity.value = 1
  selectedPaymentMethod.value = 'card'
  isProcessing.value = false
  cardValid.value = false
  cardComplete.value = false
}

function cleanupStripeElements() {
  if (cardElement.value) {
    try {
      cardElement.value.off('change')
      cardElement.value.off('ready')
      cardElement.value.off('focus')
      cardElement.value.off('blur')
      cardElement.value.unmount()
      cardElement.value.destroy()
    } catch (error) {
      console.warn('Error cleaning up Stripe element:', error)
    } finally {
      cardElement.value = null
      cardElementMounted.value = false
      cardValid.value = false
      cardComplete.value = false
    }
  }
}

async function selectPaymentMethod(methodId: string) {
  selectedPaymentMethod.value = methodId

  if (methodId === 'card' && props.isOpen) {
    await nextTick()
    if (!cardElementMounted.value) {
      await setupCardElement()
    }
  }
}

function increaseQuantity() {
  if (!props.selectedTicket) return

  const validation = validateTicketPurchase(props.selectedTicket, quantity.value + 1)
  if (validation.valid) {
    quantity.value++
  }
}

function decreaseQuantity() {
  if (quantity.value > 1) {
    quantity.value--
  }
}

async function initializeStripe() {
  if (stripeInitialized.value) return

  try {
    stripe.value = await loadStripe()
    stripeInitialized.value = true

    // Setup card element if modal is open and card method is selected
    if (props.isOpen && selectedPaymentMethod.value === 'card') {
      await nextTick()
      await setupCardElement()
    }
  } catch (error) {
    console.error('Failed to load Stripe:', error)
  }
}

async function setupCardElement() {
  if (!stripe.value || cardElementMounted.value) return

  try {
    // Clean up existing element first
    cleanupStripeElements()

    // Wait for DOM to be ready
    await nextTick()

    const cardContainer = document.getElementById('card-element')
    if (!cardContainer) {
      console.warn('Card element container not found')
      return
    }

    // Ensure container is empty before mounting
    cardContainer.innerHTML = ''

    const elements = stripe.value.elements({
      appearance: {
        theme: 'stripe',
        variables: {
          colorPrimary: '#7c3aed',
          colorBackground: '#ffffff',
          colorText: '#374151',
          colorDanger: '#ef4444',
          fontFamily: 'system-ui, sans-serif',
          spacingUnit: '4px',
          borderRadius: '8px',
        }
      }
    })

    cardElement.value = elements.create('card', {
      style: {
        base: {
          fontSize: '16px',
          color: '#374151',
          '::placeholder': {
            color: '#9ca3af',
          },
        },
      },
    })

    // Mount the element
    cardElement.value.mount(cardContainer)

    // Set up event listeners
    cardElement.value.on('change', (event: any) => {
      // Update card validity and completeness
      cardValid.value = !event.error && event.complete
      cardComplete.value = event.complete

      const displayError = document.getElementById('card-errors')
      if (displayError) {
        if (event.error) {
          displayError.textContent = event.error.message
        } else {
          displayError.textContent = ''
        }
      }
    })

    cardElement.value.on('ready', () => {
      cardElementMounted.value = true
      // console.log('Stripe card element is ready')
    })

    // cardElement.value.on('focus', () => {
    //   console.log('Card element focused')
    // })

    // cardElement.value.on('blur', () => {
    //   console.log('Card element blurred')
    // })

  } catch (error) {
    console.error('Failed to setup card element:', error)
    cardElementMounted.value = false
  }
}

async function processPayment() {
  if (!canProceedPayment.value || !props.selectedTicket || !props.user) return

  isProcessing.value = true

  try {
    if (selectedPaymentMethod.value === 'card') {
      await processCardPayment()
    } else {
      throw new Error('Payment method not implemented yet')
    }
  } catch (error: any) {
    console.error('Payment failed:', error)
    alert(`Payment failed: ${error.message}`)
  } finally {
    isProcessing.value = false
  }
}

async function processCardPayment() {
  if (!stripe.value || !cardElement.value || !props.selectedTicket || !props.user) {
    throw new Error('Payment system not ready')
  }

  // Verify card element is still mounted and valid
  if (!cardElementMounted.value) {
    throw new Error('Payment form not ready. Please try again.')
  }

  // First, create a payment method to validate the card without creating payment intent
  const { error: paymentMethodError, paymentMethod } = await stripe.value.createPaymentMethod({
    type: 'card',
    card: cardElement.value,
    billing_details: {
      name: customerInfo.value.name,
      email: customerInfo.value.email,
    },
  })

  if (paymentMethodError) {
    throw new Error(paymentMethodError.message)
  }

  if (!paymentMethod) {
    throw new Error('Failed to create payment method')
  }

  // Only create payment intent after card validation succeeds
  const paymentData = {
    ticket_id: props.selectedTicket.id,
    quantity: quantity.value
  }

  const response = await paymentStore.createPaymentIntent(paymentData)
  if (!response) {
    throw new Error('Failed to create payment intent')
  }

  // Confirm payment with the validated payment method
  const { error, paymentIntent } = await stripe.value.confirmCardPayment(
    response.client_secret,
    {
      payment_method: paymentMethod.id
    }
  )

  if (error) {
    throw new Error(error.message)
  }

  if (paymentIntent && paymentIntent.status === 'succeeded') {
    emit('success', paymentIntent.id)
    closeModal()
  } else {
    throw new Error('Payment was not completed successfully')
  }
}

// Watch for modal open/close
watch(() => props.isOpen, async (isOpen) => {
  if (isOpen) {
    // Reset form state
    resetForm()

    // Initialize Stripe if not already done
    if (!stripeInitialized.value) {
      await initializeStripe()
    }

    // Setup card element if card method is selected
    if (selectedPaymentMethod.value === 'card') {
      // Add small delay to ensure DOM is ready
      setTimeout(async () => {
        await setupCardElement()
      }, 100)
    }
  } else {
    // Clean up when modal closes
    cleanupStripeElements()
  }
})

// Watch for payment method changes
watch(selectedPaymentMethod, async (newMethod, oldMethod) => {
  if (props.isOpen) {
    if (oldMethod === 'card') {
      cleanupStripeElements()
    }

    if (newMethod === 'card' && stripe.value) {
      // Add small delay to ensure DOM is ready
      setTimeout(async () => {
        await setupCardElement()
      }, 100)
    }
  }
})

onMounted(() => {
  initializeStripe()
})

onUnmounted(() => {
  cleanupStripeElements()
})
</script>