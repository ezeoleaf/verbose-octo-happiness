export const responseSuccess = (data) => {
    return {
        success: true,
        data: data,
        message: null
    }
}

export const responseError = (msj) => {
    return {
        success: false,
        data: null,
        message: msj
    }
}