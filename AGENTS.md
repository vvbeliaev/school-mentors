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
- Frontend: Single Page Application (SPA) or SSR Monolith hosted alongside PocketBase.
- Payments: Stripe (Checkout Sessions API).
- Video: P2P WebRTC (or simple external link sharing for MVP) to minimize hosting costs.

### Domain Model & Data Schema (PocketBase Collections)

The application is modeled around the Booking Lifecycle.

#### 1. Users & Auth (users)

- Standard PocketBase auth.
- Role: ['mentor', 'mentee', 'admin']

#### 2. Profiles (profiles)

- Relation: 1:1 with users.
- Data:
  - display_name, avatar.
  - bio (Markdown support).
  - Mentor Specifics: university (e.g., Harvard, MIT), major, graduation_year, hourly_rate.
  - Mentee Specifics: current_grade, interests.

#### 3. Services/Tags (services)

- Taxonomy for searchability (e.g., "SAT Prep", "College Application", "Calculus", "Career Advice").
- Mentors link their profiles to these services.

#### 4. Availability (availability)

- Time slots defined by Mentors.
- _Simplification for MVP:_ Mentors set general "Open Hours" or ad-hoc slots.

#### 5. Bookings (bookings)

- The Aggregate Root of the business logic.
- Fields:
  - mentor_id, mentee_id.
  - scheduled_time.
  - status: ['pending_payment', 'confirmed', 'completed', 'cancelled', 'disputed'].
  - payment_id (Stripe Session ID).
  - meeting_link (URL for the video call).

---

## Core Features & Workflows (MVP Scope)

### 1. Discovery & Search (The "Marketplace")

- Goal: Mentee finds a Mentor.
- Logic: Filter by university, major, or service tag.
- UI: Grid view of Mentor cards with price/hour and credentials.

### 2. The Booking Loop (The "Transaction")

1.  Selection: Mentee views Mentor profile -> Selects a time slot.
2.  Intent: Mentee clicks "Book". System creates a booking record with status pending_payment.
3.  Payment: System redirects to Stripe Checkout (Client requirement: One-time payment, no subscriptions yet).
4.  Confirmation: Stripe Webhook -> Updates booking status to confirmed.
5.  Notification: Both parties receive a confirmation (Email/In-app notification).

### 3. The Session (The "Product")

- Video Strategy:
  - _Constraint:_ Client is worried about video hosting costs.
  - _Solution:_ Use browser-native P2P WebRTC (e.g., simple-peer) OR allow Mentor to paste a Google Meet/Zoom link into the booking details upon confirmation.
  - MVP Decision: Simplicity first. Platform generates a unique "Meeting Room" URL where peers connect directly.

### 4. Review System (The "Network Effect")

- Post-booking, Mentees can leave a 1-5 star rating and text review.
- _Critical for trust:_ New users need to see social proof to book.

---

## Business Logic & Constraints

### Pricing Model

- Current: Pay-per-booking.
- Future: In-app currency ("Credits") to encourage bulk buying (e.g., "Buy 5 sessions, get 10% off").
- Take Rate: Platform adds a % fee on top of the Mentor's asking price OR deducts from the payout.

### Market Validation (Risks)

- Cold Start Problem: The app is useless without a critical mass of Mentors.
- MVP Focus: The prototype must feel complete enough to onboard the first batch of university students manually.
- Pivot Resistance: As noted by the developer, the model relies heavily on network effects. If user volume is low, the app offers little value. The UI should prioritize "Featured Mentors" to make the platform look busy even with few users.

---

## Development Priorities (Agent Instructions)

1.  Scaffold PocketBase: Set up collections with strict API rules (ACLs) to ensure Mentees cannot edit Mentor profiles.
2.  Auth Flow: Implement Email/Password flow tailored for two distinct personas.
3.  Payment Integration: Build a minimal Stripe wrapper to handle the pending -> paid state transition securely.
4.  UI/UX: Focus on the Mentor Profile page—this is the primary sales page for the application. It needs to look professional to justify the booking fee.

---

### Client Communication Context

- Client Persona: Non-technical founder. Overestimates speed ("take a few days to build") and underestimates complexity.
- Strategy: Deliver a robust MVP (Prototype) that handles the "Happy Path" (Search -> Book -> Pay -> Meet) flawlessly, ignoring complex edge cases (like refunds, rescheduling) for version 1.

## Business rules

1.  Booking Logic: A Mentee cannot book a slot that is already booked. Need atomic transactions or strict checks in PocketBase API rules.
2.  Identity: Two distinct roles (mentor, mentee). A Mentee primarily _reads_ profiles and _writes_ bookings. A Mentor _writes_ profiles and slots.
3.  Payment Gate: The service is useless without payment gating. No meeting link reveals until booking.status == 'paid'.
4.  MVP Shortcuts:
    - No automated payouts to mentors (Manual wire transfer at end of month).
    - No in-app messaging (Email notifications only).
    - No recurring subscriptions.

### 1. Unified Identity

- Concept: No strict separation between Mentor and Mentee accounts. Every user is a Mentee by default. Any user can _become_ a Mentor by completing a profile and enabling availability.
- Architecture:
  - One User, One Profile.
  - A boolean flag is_listing_active determines if they appear in search results.

### 2. Search Entity: The Person (Not the Gig)

- Decision: We are an "Expert Network," not a "Service Gig Market".
- Implementation:
  - Users verify their identity/university once.
  - They define a list of topics or tags they can help with (e.g., #Calculus, #LifeCoaching).
  - Search Logic: Users search for _topics_, but the results display _People Cards_ (Profiles).
  - Pricing: MVP simplification -> Single hourly rate per Mentor (regardless of the topic).

### 3. Data Schema Adjustments (PocketBase)

#### users (System)

- id, email, password

- name, avatar
- university, degree, graduation_year
- bio (Markdown)
- hourlyRate (Int)
- topics (JSON array or Relation to tags collection) — _Critical for search_
- isMentor (Bool) — \_Index this field\*

#### bookings

- Links mentor_id (User A) and mentee_id (User B).
- _Note:_ User A can be a mentor in this transaction, but a mentee in another transaction tomorrow.

## Booking & Scheduling Architecture (MVP)

### Strategy: Explicit Slots

Instead of recurring schedules, Mentors explicitly create "Availability Slots" (Time Blocks).

### Schema Definition

1. **slots**
   - Represents a specific time block offered by a Mentor.
   - Fields:
     - mentor_id (Rel: users)
     - start_time (DateTime, UTC)
     - is_booked (Boolean) - Acts as a lock.

2. **bookings**
   - Represents the transaction/contract between peers.
   - Fields:
     - mentee_id (Rel: users)
     - slot_id (Rel: availability_slots, Unique/1-to-1)
     - status: pending, confirmed, completed
     - meeting_url: Generated post-payment.

### Meeting Logic (VProvider:vider:** Jitsi Meet (Public public instaLink Generation:ation:** Deterministic URL pattern: https://meet.jit.si/AppNamespace_{booking_id}.

- **Access:** Link is revealed in UI only when booking.status === 'confirmed'.

### TimezoneDatabase:\*Database:\*\* Always stores start_time in UTC.

- **Frontend:** Must convert UTC -> Local Time for rendering calendars.
