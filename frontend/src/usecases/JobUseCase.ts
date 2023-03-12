import {
  JobRepository,
  JobRepositoryImpl,
} from "../repositories/JobRepository"
import Job from "../models/Job"
import { ENDPOINT_URL } from "../configs/config"
import SignedUrl from "../models/SignedUrl"
import SignedUrlRepositoryImpl, {
  SignedUrlRepository,
} from "../repositories/SignedUrlRepository"

export interface JobUseCase {
  startJob(file: File): Promise<boolean>
  getJobs(): Promise<Job[] | undefined>
  getJob(id: string): Promise<Job | undefined>
  deleteJob(id: string): Promise<boolean>
  getSignedUrl(key: string): Promise<SignedUrl | undefined>
}

export default class JobUseCaseImpl implements JobUseCase {
  private jobRepository: JobRepository
  private signedUrlRepository: SignedUrlRepository

  constructor() {
    this.jobRepository = new JobRepositoryImpl(
      `${ENDPOINT_URL}/detect_texts`
    )
    this.signedUrlRepository = new SignedUrlRepositoryImpl(
      `${ENDPOINT_URL}/signed_urls`
    )
  }

  getJob = async (id: string) => {
    try {
      return await this.jobRepository.getById({ id })
    } catch (e) {
      console.error(e)
    }
    return undefined
  }

  getJobs = async () => {
    try {
      return await this.jobRepository.getAll()
    } catch (e) {
      console.error(e)
    }
    return undefined
  }

  startJob = async (file: File) => {
    try {
      const result = await this.jobRepository.create({ file })
      return result.status
    } catch (e) {
      console.error(e)
    }
    return false
  }

  deleteJob = async (id: string) => {
    try {
      const result = await this.jobRepository.delete({ id })
      return result.status
    } catch (e) {
      console.error(e)
    }
    return false
  }

  getSignedUrl = async (key: string) => {
    try {
      return await this.signedUrlRepository.get({ key })
    } catch (e) {
      console.error(e)
    }
    return undefined
  }
}
