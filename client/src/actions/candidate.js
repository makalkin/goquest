import {CANDIDATE} from '../constants/actionTypes';

export const addCandidate = (candidate) => {
	return {
		type: CANDIDATE.ADD_CANDIDATE,
		payload: {
			request: {
				url: "/api/v1/Candidates",
				method: "post",
				data: {
					...candidate
				}
			}
		}
	}
};


export const getCandidate = (filters) => {
	return {
		type: CANDIDATE.GET_CANDIDATE,
		payload: {
			request: {
				url: "/api/v1/Candidates",
				method: "get",
				params: {
					...filters
				}
			}
		}
	}
};