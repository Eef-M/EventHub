<template>
  <Card>
    <CardContent class="py-2 px-4">
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
              <Button size="icon" class="bg-slate-600 hover:bg-slate-700 cursor-pointer" @click="$emit('edit', ticket)">
                <Pencil class="w-4 h-4" />
              </Button>
              <Button size="icon" class="bg-red-600 hover:bg-red-700 cursor-pointer" @click="$emit('delete', ticket)">
                <Trash class="w-4 h-4" />
              </Button>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </CardContent>
  </Card>
</template>

<script setup lang="ts">
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
import { Pencil, Trash } from 'lucide-vue-next'
import type { OrganizerTicket } from "@/types/organizer";

interface Props {
  tickets: OrganizerTicket[]
}

defineProps<Props>()

defineEmits<{
  edit: [ticket: OrganizerTicket]
  delete: [ticket: OrganizerTicket]
}>()
</script>