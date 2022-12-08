import { get, writable } from 'svelte/store';
import type Region from '$lib/types/Region';
import type Candidate from '$lib/types/Candidate';
import { TossupCandidateStore } from './Candidates';

const EditCandidateModalStore = writable<{
	candidate: Candidate;
	open: boolean;
}>({
	candidate: get(TossupCandidateStore),
	open: false
});

const ClearMapModalStore = writable({
	open: false
});

const MapModalStore = writable({
	open: false
});

const EditRegionModalStore = writable<{
	region: Region | null;
	open: boolean;
}>({
	region: null,
	open: false
});

export { EditCandidateModalStore, ClearMapModalStore, MapModalStore, EditRegionModalStore };