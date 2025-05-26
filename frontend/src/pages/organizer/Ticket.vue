<template>
  <OrganizerLayout>
    <div class="p-6 space-y-6">
      <div class="flex justify-between items-center">
        <h2 class="text-2xl font-semibold">Manage Tickets</h2>
        <Button class="bg-purple-600 hover:bg-purple-700 cursor-pointer" @click="openCreateModal">
          <Plus class="w-4 h-4 mr-2" /> Create Ticket
        </Button>
      </div>

      <Card>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Ticket Name</TableHead>
                <TableHead>Price</TableHead>
                <TableHead>Quota</TableHead>
                <TableHead>Event</TableHead>
                <TableHead class="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-if="!tickets.length && !tickets">
                <TableCell :colspan="5" class="text-center text-muted-foreground">
                  No tickets available
                </TableCell>
              </TableRow>
              <TableRow v-for="ticket in tickets" :key="ticket.id" class="hover:bg-muted transition">
                <TableCell>{{ ticket.name }}</TableCell>
                <TableCell>$ {{ ticket.price.toLocaleString() }}</TableCell>
                <TableCell>{{ ticket.quota }}</TableCell>
                <TableCell>{{ ticket.event_title }}</TableCell>
                <TableCell class="text-right space-x-2">
                  <Button size="icon" class="bg-slate-600 hover:bg-slate-700 cursor-pointer"
                    @click="openEditModal(ticket)">
                    <Pencil class="w-4 h-4" />
                  </Button>
                  <Button size="icon" class="bg-red-600 hover:bg-red-700 cursor-pointer"
                    @click="openDeleteModal(ticket)">
                    <Trash class="w-4 h-4" />
                  </Button>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </CardContent>
      </Card>

      <div v-if="tickets.length === 0" class="text-gray-500">
        No tickets available.
      </div>

      <!-- Modal: Create/Edit -->
      <Dialog :open="showFormModal" @update:open="showFormModal = $event">
        <DialogContent class="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>{{ isEditing ? "Edit Ticket" : "Add Ticket" }}</DialogTitle>
          </DialogHeader>
          <div class="grid gap-4 py-4">
            <div class="grid gap-2">
              <Label for="name">Ticket Name</Label>
              <Input v-model="form.name" id="name" placeholder="e.g. VIP" />
            </div>
            <div class="grid gap-2">
              <Label for="description">Description</Label>
              <Input v-model="form.description" id="description" placeholder="desc..." />
            </div>
            <div class="grid gap-2">
              <Label for="price">Price</Label>
              <Input v-model.number="form.price" id="price" type="number" />
            </div>
            <div class="grid gap-2">
              <Label for="quota">Quota</Label>
              <Input v-model.number="form.quota" id="quota" type="number" />
            </div>
            <div class="grid gap-2">
              <Label for="event">Event</Label>
              <Select v-model="form.event_id">
                <SelectTrigger id="event" class="w-full">
                  <SelectValue placeholder="Select an event" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem v-for="event in events" :key="event.id" :value="event.id">
                    {{ event.title }}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" @click="showFormModal = false" class="cursor-pointer">Cancel</Button>
            <Button
              :class="isEditing ? 'bg-slate-600 hover:bg-slate-700 cursor-pointer' : 'bg-purple-600 hover:bg-purple-700 cursor-pointer'"
              @click="handleSubmitTicket">
              {{ isEditing ? "Update" : "Create" }}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      <!-- Modal: Confirm Delete -->
      <Dialog :open="showDeleteModal" @update:open="showDeleteModal = $event">
        <DialogContent class="sm:max-w-sm">
          <DialogHeader>
            <DialogTitle>Delete Ticket</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete <strong>{{ selectedTicket?.name }}</strong>?
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button class="cursor-pointer" variant="outline" @click="showDeleteModal = false">Cancel</Button>
            <Button class="bg-red-600 hover:bg-red-700 cursor-pointer" @click="deleteTicket">Delete</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  </OrganizerLayout>
</template>

<script setup lang="ts">
import OrganizerLayout from "@/layouts/OrganizerLayout.vue";
import { computed, onMounted, reactive, ref } from "vue";
import {
  Card,
  CardContent,
} from "@/components/ui/card";
import {
  Table,
  TableHeader,
  TableBody,
  TableRow,
  TableHead,
  TableCell,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Pencil, Plus, Trash } from "lucide-vue-next";
import { toast } from "vue-sonner";
import { useOrganizerStore } from "@/stores/organizerStore";
import { useTicketStore } from "@/stores/ticketStore";

const organizerStore = useOrganizerStore()
const ticketStore = useTicketStore()
const tickets = computed(() => organizerStore.ticketsState.data)
const events = computed(() => organizerStore.eventsState.data)

onMounted(() => {
  organizerStore.getOrganizerTickets()
  organizerStore.getMyEvents()
})

const showFormModal = ref(false);
const showDeleteModal = ref(false);
const isEditing = ref(false);
const selectedTicket = ref<any | null>(null);

const defaultForm = {
  event_id: "",
  name: "",
  description: "",
  price: 0,
  quota: 0,
};

const form = reactive({ ...defaultForm });

function openCreateModal() {
  isEditing.value = false;
  Object.assign(form, defaultForm);
  showFormModal.value = true;
}

function openEditModal(ticket: any) {
  isEditing.value = true;
  Object.assign(form, {
    event_id: ticket.event_id,
    name: ticket.name,
    description: ticket.description,
    price: ticket.price,
    quota: ticket.quota,
  });
  selectedTicket.value = ticket;
  showFormModal.value = true;
}

async function handleSubmitTicket() {
  try {
    const formData = new FormData();
    formData.append("event_id", form.event_id);
    formData.append("name", form.name);
    formData.append("description", form.description);
    formData.append("price", form.price.toString());
    formData.append("quota", form.quota.toString());

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
    Object.assign(form, defaultForm);
  } catch (err) {
    toast.error(`Failed to ${isEditing.value ? 'update' : 'create'} ticket`, {
      description: `An unexpected error occurred. ${isEditing.value ? ticketStore.updateState.error : ticketStore.createState.error}`,
    })
    console.error(`Failed to ${isEditing.value ? 'update' : 'create'} ticket`, err);
  }
}

function openDeleteModal(ticket: any) {
  selectedTicket.value = ticket;
  showDeleteModal.value = true;
}

function deleteTicket() {
  toast.success('Ticket deleted successfully', {
    description: new Date().toLocaleString(),
  })
  showDeleteModal.value = false;
}
</script>
