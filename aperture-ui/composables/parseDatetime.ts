function toLocale(dateTimeObject: Date) {
    return dateTimeObject.toLocaleString().split(',')[0];
}

export const parseDatetime = (startDate: string, endDate: string) => {
    const startDateDTO = new Date(startDate);
    const endDateDTO = new Date(endDate);
    if (startDate == endDate){
        return toLocale(startDateDTO);
    }
    else {

        return toLocale(startDateDTO) + " - " + toLocale(endDateDTO)
    }
    
}

export const albumDatetime = (startDate: string) => {
    const startDateDTO = new Date(startDate);
    const month = startDateDTO.toLocaleString('default', { month: 'long' });
    return month + " " + startDateDTO.getFullYear()
}