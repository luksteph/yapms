import assert from 'assert';
import { SavedMapSchema } from '../lib/types/SavedMap';

describe('array', function () {
	const data = {};
	const parsedData = SavedMapSchema.safeParse(data);
	assert(parsedData.success === true);
});
