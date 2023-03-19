
import axiosOCR from '../utils/axios/axiosOCR';

import authHeader from './auth-header';

class OCRService {
  postImageData(data) {
    console.log(authHeader())
    console.log('dsadas')
    return axiosOCR.post('/', data, {headers: authHeader()});
  }
}

export default new OCRService();