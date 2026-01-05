import PocketBase from 'pocketbase';
import { nanoid } from 'nanoid';

const UNIVERSITIES = [
	'Harvard University',
	'Stanford University',
	'MIT',
	'Oxford University',
	'UC Berkeley',
	'University of Toronto',
	'ETH Zurich',
	'Cambridge University',
	'Princeton University',
	'Yale University'
];

const DEGREES = [
	'Computer Science',
	'Economics',
	'Physics',
	'Mechanical Engineering',
	'Mathematics',
	'Business Administration',
	'Biology',
	'Architecture',
	'Political Science',
	'Philosophy'
];

const TAGS_POOL = [
	'SAT Prep',
	'Calculus',
	'College Essays',
	'Internship Advice',
	'Coding',
	'Machine Learning',
	'Resume Review',
	'Interview Prep',
	'Data Science',
	'Product Management',
	'Quant Finance',
	'Medical School Apps',
	'Law School Apps'
];

const NAMES = [
	'Alex Rivera',
	'Jordan Chen',
	'Sarah Miller',
	'David Kumar',
	'Elena Popova',
	'James Wilson',
	'Maya Gupta',
	'Lucas Schmidt',
	'Aisha Bello',
	"Liam O'Connor",
	'Sofia Rodriguez',
	'Noah Williams',
	'Emma Brown',
	'Ethan Jones',
	'Olivia Garcia',
	'William Davis',
	'Isabella Martinez',
	'Mason Hernandez',
	'Mia Lopez',
	'Benjamin Gonzalez'
];

const NUM_MENTORS = 100;

async function seed() {
	const [url, email, password] = process.argv.slice(2);

	if (!url || !email || !password) {
		console.error('Usage: npx tsx scripts/seed-mentors.ts <pb_url> <admin_email> <admin_password>');
		process.exit(1);
	}

	const pb = new PocketBase(url);

	try {
		console.log(`Authenticating as admin ${email}...`);
		await pb.admins.authWithPassword(email, password);
		console.log('Authenticated successfully.');
	} catch (err) {
		console.error('Failed to authenticate as admin:', err.message);
		process.exit(1);
	}

	console.log(`Starting to seed ${NUM_MENTORS} mentors...`);

	for (let i = 0; i < NUM_MENTORS; i++) {
		const name = NAMES[i % NAMES.length];
		const university = UNIVERSITIES[Math.floor(Math.random() * UNIVERSITIES.length)];
		const degree = DEGREES[Math.floor(Math.random() * DEGREES.length)];
		const graduationYear = 2020 + Math.floor(Math.random() * 8);
		const hourRateCents = (20 + Math.floor(Math.random() * 80)) * 100;

		// Pick 2-4 random tags
		const numTags = 2 + Math.floor(Math.random() * 3);
		const tags = [...TAGS_POOL].sort(() => 0.5 - Math.random()).slice(0, numTags);

		const emailPrefix = name.toLowerCase().replace(/\s+/g, '.');
		const userEmail = `${emailPrefix}.${nanoid(4)}@example.com`;
		const userPassword = 'Password123!';

		const data = {
			username: `${emailPrefix}_${nanoid(4)}`,
			email: userEmail,
			emailVisibility: true,
			password: userPassword,
			passwordConfirm: userPassword,
			name: name,
			university: university,
			degree: degree,
			graduationYear: graduationYear,
			bio: `Hi! I am ${name}, a ${degree} student at ${university}. I love helping others achieve their academic goals and navigating career paths in ${tags.slice(0, 2).join(' and ')}. Feel free to book a session with me!`,
			tags: tags,
			isMentor: true,
			isVerified: true,
			hourRateCents: hourRateCents
		};

		try {
			const record = await pb.collection('users').create(data);
			console.log(
				`[${i + 1}/${NUM_MENTORS}] Created mentor: ${record.name} (${record.email}) at ${record.university}`
			);
		} catch (err) {
			console.error(
				`[${i + 1}/${NUM_MENTORS}] Failed to create mentor ${name}:`,
				err.data || err.message
			);
		}
	}

	console.log('Seeding completed.');
}

seed();
