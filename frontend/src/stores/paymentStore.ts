import { PaymentService } from "@/services/paymentService";
import type { CreatePaymentRequest, CreatePaymentResponse, Payment, PaymentHistoryResponse, PaymentStatus } from "@/types/payment";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

interface PaymentState {
  data: Payment[] | null;
  loading: boolean;
  error: string | null;
}

interface CreatePaymentState {
  loading: boolean;
  error: string | null;
  clientSecret: string | null;
}

export const usePaymentStore = defineStore('payment', () => {
  // State
  const paymentsState = ref<PaymentState>({
    data: null,
    loading: false,
    error: null
  });

  const createPaymentState = ref<CreatePaymentState>({
    loading: false,
    error: null,
    clientSecret: null
  });

  const currentPayment = ref<Payment | null>(null);
  const pagination = ref({
    page: 1,
    limit: 10,
    total: 0
  });

  // Getters
  const payments = computed(() => paymentsState.value.data);
  const isPaymentsLoading = computed(() => paymentsState.value.loading);
  const paymentsError = computed(() => paymentsState.value.error);

  const isCreatingPayment = computed(() => createPaymentState.value.loading);
  const createPaymentError = computed(() => createPaymentState.value.error);
  const clientSecret = computed(() => createPaymentState.value.clientSecret);

  const successfulPayments = computed(() =>
    payments.value?.filter(payment => payment.status === 'succeeded') || []
  );

  const pendingPayments = computed(() =>
    payments.value?.filter(payment => payment.status === 'pending') || []
  );

  // Actions
  async function createPaymentIntent(data: CreatePaymentRequest): Promise<CreatePaymentResponse | null> {
    createPaymentState.value.loading = true;
    createPaymentState.value.error = null;
    createPaymentState.value.clientSecret = null;

    try {
      const response = await PaymentService.createPaymentIntent(data);
      createPaymentState.value.clientSecret = response.client_secret;
      return response;
    } catch (error: any) {
      createPaymentState.value.error = error.message;
      throw error;
    } finally {
      createPaymentState.value.loading = false;
    }
  }

  async function getPaymentHistory(page = 1, limit = 10): Promise<void> {
    paymentsState.value.loading = true;
    paymentsState.value.error = null;

    try {
      const response: PaymentHistoryResponse = await PaymentService.getPaymentHistory(page, limit);

      paymentsState.value.data = response.payments;
      pagination.value = {
        page: response.page,
        limit: response.limit,
        total: response.total
      };
    } catch (error: any) {
      paymentsState.value.error = error.message;
    } finally {
      paymentsState.value.loading = false;
    }
  }

  async function getPayment(paymentId: string): Promise<Payment | null> {
    try {
      const payment = await PaymentService.getPayment(paymentId);
      currentPayment.value = payment;
      return payment;
    } catch (error: any) {
      console.error('Error getting payment:', error);
      return null;
    }
  }

  function updatePaymentStatus(paymentId: string, status: PaymentStatus): void {
    if (paymentsState.value.data) {
      const paymentIndex = paymentsState.value.data.findIndex(p => p.id === paymentId);
      if (paymentIndex !== -1) {
        paymentsState.value.data[paymentIndex].status = status;
      }
    }

    if (currentPayment.value?.id === paymentId) {
      currentPayment.value.status = status;
    }
  }

  function clearCreatePaymentState(): void {
    createPaymentState.value = {
      loading: false,
      error: null,
      clientSecret: null
    };
  }

  function clearPaymentsState(): void {
    paymentsState.value = {
      data: null,
      loading: false,
      error: null
    };
    currentPayment.value = null;
    pagination.value = {
      page: 1,
      limit: 10,
      total: 0
    };
  }

  return {
    // State
    paymentsState,
    createPaymentState,
    currentPayment,
    pagination,

    // Getters
    payments,
    isPaymentsLoading,
    paymentsError,
    isCreatingPayment,
    createPaymentError,
    clientSecret,
    successfulPayments,
    pendingPayments,

    // Actions
    createPaymentIntent,
    getPaymentHistory,
    getPayment,
    updatePaymentStatus,
    clearCreatePaymentState,
    clearPaymentsState
  };
})