<template>
  <OrganizerLayout>
    <div class="p-6 space-y-6">
      <div class="flex justify-between items-center">
        <h2 class="text-2xl font-semibold">Manage Tickets</h2>
        <Button class="bg-purple-600 hover:bg-purple-700 cursor-pointer" @click="openCreateModal">
          <Plus class="w-4 h-4 mr-2" /> Create Ticket
        </Button>
      </div>

      <TicketsTable :tickets="organizerStore.ticketsState.data" @edit="openEditModal" @delete="openDeleteModal" />

      <TicketFormModal v-model:isOpen="showFormModal" :isEditing="isEditing" :ticketData="selectedTicket"
        :events="organizerStore.eventsState.data"
        :isLoading="isEditing ? ticketStore.updateState.loading : ticketStore.createState.loading"
        @submit="handleSubmitTicket" @cancel="handleCancelForm" />

      <DeleteConfirmModal v-model:isOpen="showDeleteModal" :ticketName="selectedTicket?.name"
        :isLoading="ticketStore.deleteState.loading" @confirm="confirmDelteTicket" @cancel="handleCancelDelete" />

    </div>
  </OrganizerLayout>
</template>

<script setup lang="ts">
import OrganizerLayout from "@/layouts/OrganizerLayout.vue";
import TicketsTable from "@/components/organizer/tickets/TicketsTable.vue";
import TicketFormModal from "@/components/organizer/tickets/TicketFormModal.vue";
import DeleteConfirmModal from "@/components/organizer/tickets/DeleteConfirmModal.vue";
import { onMounted, ref } from "vue";
import { Plus } from "lucide-vue-next";
import { toast } from "vue-sonner";
import { useOrganizerStore } from "@/stores/organizerStore";
import { useTicketStore } from "@/stores/ticketStore";
import type { Ticket } from "@/types/ticket";
import { Button } from "@/components/ui/button";

const organizerStore = useOrganizerStore()
const ticketStore = useTicketStore()

const showFormModal = ref(false);
const showDeleteModal = ref(false);
const isEditing = ref(false);
const selectedTicket = ref<Ticket | null>(null);

onMounted(() => {
  organizerStore.getOrganizerTickets()
  organizerStore.getMyEvents()
})

function openCreateModal() {
  isEditing.value = false;
  selectedTicket.value = null
  showFormModal.value = true;
}

function openEditModal(ticket: Ticket) {
  isEditing.value = true;
  selectedTicket.value = ticket;
  showFormModal.value = true;
}

function openDeleteModal(ticket: Ticket) {
  selectedTicket.value = ticket;
  showDeleteModal.value = true;
}

async function handleSubmitTicket(formData: any) {
  try {
    const formDataToSend = new FormData();
    formDataToSend.append("event_id", formData.event_id);
    formDataToSend.append("name", formData.name);
    formDataToSend.append("description", formData.description);
    formDataToSend.append("price", formData.price.toString());
    formDataToSend.append("quota", formData.quota.toString());

    if (isEditing.value && selectedTicket.value) {
      await ticketStore.updateTicket(selectedTicket.value.id, formData)
    } else {
      await ticketStore.createTicket(formData)
    }

    await organizerStore.getOrganizerTickets()

    toast.success(`Ticket ${isEditing.value ? 'updated' : 'created'} successfuly`, {
      description: new Date().toLocaleString(),
    })
    showFormModal.value = false;
    selectedTicket.value = null
  } catch (err) {
    toast.error(`Failed to ${isEditing.value ? 'update' : 'create'} ticket`, {
      description: `An unexpected error occurred. ${isEditing.value ? ticketStore.updateState.error : ticketStore.createState.error}`,
    })
    console.error(`Failed to ${isEditing.value ? 'update' : 'create'} ticket`, err);
  }
}

async function confirmDelteTicket() {
  try {
    if (!selectedTicket.value) return;

    await ticketStore.deleteTicket(selectedTicket.value.id)
    await organizerStore.getOrganizerTickets()

    toast.success('Ticket deleted successfully', {
      description: new Date().toLocaleString(),
    })

    showDeleteModal.value = false;
    selectedTicket.value = null;
  } catch (err) {
    toast.error('Failed to delete ticket', {
      description: ticketStore.deleteState.error || `An unexpected error occurred. ${err}`,
    })
    console.error('Delete Ticket Failed: ', ticketStore.deleteState.error, err)
  }
}

function handleCancelForm() {
  selectedTicket.value = null
}

function handleCancelDelete() {
  selectedTicket.value = null
}
</script>