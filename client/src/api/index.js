/**
 * The purpose of this module is to collect all usable APIs under one access point (<root>/src/api).
 *
 * This approach prevents duplication of API class objects yet doesn't restrict import of APIs which may be useful
 * for setting custom configurations for APIs in the future.
 */

import AuthAPI from './auth';
import CandidateAPI from './candidate';
import FileAPI from './file';
import ContactInfoAPI from './contactInfo';
import LocalizationAPI from './localization';
import RecruiterAPI from './recruiter';
import UserAPI from './user';
import VacancyAPI from './vacancy';
import AccountAPI from './account';


export const auth = new AuthAPI();
export const candidate = new CandidateAPI();
export const file = new FileAPI();
export const contactInfo = new ContactInfoAPI();
export const localization = new LocalizationAPI();
export const recruiter = new RecruiterAPI();
export const user = new UserAPI();
export const vacancy = new VacancyAPI();
export const account = new AccountAPI();
