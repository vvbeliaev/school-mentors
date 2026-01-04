## Project Context & Vision

Project Name: (TBD) Peer-to-Peer Student Mentorship Platform
Type: Two-sided Marketplace / EdTech
Architecture: PocketBase-driven Monolith
Status: MVP / Prototype Phase

Core Concept:
A centralized platform connecting youth students (Mentees) with current/former college students (Mentors) for personalized 1-on-1 coaching, tutoring, and advice. The goal is to democratize access to academic and career guidance by leveraging the "near-peer" model.

Business Value Proposition:

1.  For Mentees: Access to relatable, affordable guidance from someone who "just did it" (got into that college, passed that exam).
2.  For Mentors: Flexible income stream and CV-building experience.
3.  Platform: Generates revenue via transaction fees (booking commissions).

---

## Technical Architecture (The "How")

### Infrastructure Strategy

- Backend: PocketBase (Golang). Used as the all-in-one backend for:
  - Authentication (Email/Pass, potentially OAuth later).
  - Database (SQLite with WAL).
  - File Storage (Profile pictures, verification docs).
  - Real-time subscriptions (Booking updates).
  - Business Logic: Golang hooks and cron jobs for booking lifecycle.
- Frontend: SvelteKit SPA (hosted alongside PocketBase).
- Payments: Stripe (Payment Intents API).
- Video: Deterministic Jitsi Meet links.

### Domain Model & Data Schema (PocketBase Collections)

The application is built around a Unified Identity model and an explicit Slot-based scheduling system.

#### 1. Users & Auth (users)

- One User, One Profile. Every user can be a Mentee; some are also Mentors.
- Data:
  - `name`, `avatar`, `bio` (Markdown support).
  - `university`, `degree`, `graduationYear`.
  - `isMentor` (Boolean): Flag to appear in search results.
  - `hourRateCents` (Integer): Hourly rate in cents.
  - `tags` (JSON): Skills and topics (e.g., ["SAT Prep", "Calculus"]).
  - `isVerified` (Boolean): Manual verification status.
  - `stripeCustomer` (String): Reference to Stripe.

#### 2. Slots (slots)

- Represents an explicit availability block offered by a Mentor.
- Fields:
  - `mentor` (Relation: users).
  - `start` (DateTime): UTC start time.
  - `durationMinutes` (Integer): Length of the session.
  - `booked` (Boolean): Acts as a lock to prevent double-booking.

#### 3. Bookings (bookings)

- The transaction record between a Mentee and a Mentor (via a Slot).
- Fields:
  - `mentee` (Relation: users).
  - `slot` (Relation: slots, 1-to-1).
  - `status`: [`pending`, `confirmed`, `canceled`, `expired`].
  - `agreedPriceCents` (Integer): Price locked at the time of booking creation.
  - `stripePaymentIntent` (String): Link to the Stripe transaction.
  - `meta` (JSON): Additional metadata.

---

## Core Features & Workflows (MVP Scope)

### 1. Discovery & Search (The "Marketplace")

- Goal: Mentee finds a Mentor by topic or credentials.
- UI: Grid view of Mentor cards showing price, university, and tags.

### 2. The Booking Loop (The "Transaction")

1.  Selection: Mentee selects an available slot on the Mentor's profile.
2.  Intent: Mentee clicks "Book".
    - System creates a booking record with status `pending`.
    - **Hook:** Automatically marks the `slot.booked = true`.
    - **Hook:** Copies Mentor's `hourRateCents` to `booking.agreedPriceCents`.
3.  Payment: Mentee proceeds to pay via Stripe.
4.  Confirmation: Upon successful payment (via webhook or client-side redirect), booking status updates to `confirmed`.
5.  Expiration: A cron job runs every 10 minutes to expire `pending` bookings older than 15 minutes, releasing the slot.

### 3. The Session (The "Product")

- Video Strategy: Deterministic Jitsi Meet rooms.
- URL Pattern: `https://meet.jit.si/AppNamespace_{booking_id}`.
- Access: Meeting link is only revealed in the UI when `booking.status === 'confirmed'`.

---

## Business Logic & Constraints

### 1. Atomic Locking

- PocketBase hooks ensure a slot cannot be booked twice. If `slot.booked` is already true, the booking creation is rejected.

### 2. Status Lifecycle

- `pending`: Initial state, slot is reserved for 15 minutes.
- `confirmed`: Payment successful, meeting is active.
- `canceled`/`expired`: Slot is released (`slot.booked = false`).

### 3. Pricing Model

- Fixed hourly rate per Mentor (stored in cents).
- No platform fees for MVP (Mentors get 100%).

### 4. Timezones

- Backend stores all times in UTC.
- Frontend renders based on the user's local timezone.
