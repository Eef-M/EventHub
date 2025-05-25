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
                  No tickets found
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
          <div class="space-y-4">
            <div>
              <Label for="name">Ticket Name</Label>
              <Input v-model="form.name" id="name" placeholder="e.g. VIP" />
            </div>
            <div>
              <Label for="price">Price</Label>
              <Input v-model.number="form.price" id="price" type="number" />
            </div>
            <div>
              <Label for="quota">Quota</Label>
              <Input v-model.number="form.quota" id="quota" type="number" />
            </div>
            <div>
              <Label for="event">Event Title</Label>
              <Select v-model="form.event_title">
                <SelectTrigger id="event" class="w-full">
                  <SelectValue placeholder="Select an event" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem v-for="event in events" :key="event.id" :value="event.title">
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
              @click="saveTicket">
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
import { computed, onMounted, ref } from "vue";
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

const organizerStore = useOrganizerStore()
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

const form = ref({
  id: "",
  name: "",
  price: 0,
  quota: 0,
  event_title: "",
});

function openCreateModal() {
  isEditing.value = false;
  form.value = { id: "", name: "", price: 0, quota: 0, event_title: "" };
  showFormModal.value = true;
}

function openEditModal(ticket: any) {
  isEditing.value = true;
  form.value = { ...ticket };
  showFormModal.value = true;
}

function saveTicket() {
  if (isEditing.value) {
    toast.success('Ticket updated successfully', {
      description: new Date().toLocaleString(),
    })
  } else {
    toast.success('Ticket created successfully', {
      description: new Date().toLocaleString(),
    })
  }
  showFormModal.value = false;
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
