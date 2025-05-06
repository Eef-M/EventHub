// dummies data
export interface Event {
  id: string;
  title: string;
  description: string;
  location: string;
  date: string;
  category: string;
  bannerURL: string;
}

export const dummyEvents: Event[] = [
  {
    id: "1",
    title: 'Vue Conference 2025',
    description: '',
    location: 'Jakarta',
    date: '2025-06-10',
    category: 'Webinar',
    bannerURL: 'https://images.pexels.com/photos/546819/pexels-photo-546819.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
  },
  {
    id: "2",
    title: 'Tech Music Fest',
    description: '',
    location: 'Bandung',
    date: '2025-05-12',
    category: 'Concert',
    bannerURL: 'https://images.pexels.com/photos/3861969/pexels-photo-3861969.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
  },
  {
    id: "3",
    title: 'Startup Workshop',
    description: '',
    location: 'Surabaya',
    date: '2025-07-01',
    category: 'Workshop',
    bannerURL: 'https://images.pexels.com/photos/1595385/pexels-photo-1595385.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
  },
  {
    id: "4",
    title: 'Design Thinking Meetup',
    description: '',
    location: 'Yogyakarta',
    date: '2025-08-15',
    category: 'Meetup',
    bannerURL: 'https://images.pexels.com/photos/57690/pexels-photo-57690.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
  },
  {
    id: "5",
    title: 'AI & Future Talk',
    description: '',
    location: 'Bali',
    date: '2025-09-05',
    category: 'Seminar',
    bannerURL: 'https://images.pexels.com/photos/6153354/pexels-photo-6153354.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
  },
  {
    id: "6",
    title: 'Freelancer Networking Night',
    description: '',
    location: 'Jakarta',
    date: '2025-10-10',
    category: 'Networking',
    bannerURL: 'https://images.pexels.com/photos/2773507/pexels-photo-2773507.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1'
  }
];
