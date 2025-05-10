import {FormControl, InputLabel, MenuItem, Select} from "@mui/material";

export const SelectInput = ({
    label,
    setFieldValue,
    items,
    value,
    ...otherProps
}) => {
    const onChange = event => {
        setFieldValue(event.target.value)
    };

    return (
        <FormControl variant="standard" sx={otherProps}>
            <InputLabel>{label}</InputLabel>
            <Select
                value={value}
                onChange={onChange}
                {...otherProps}
            >
                {items.map(item => (
                    <MenuItem key={item} value={item}>
                        {item}
                    </MenuItem>
                ))}
            </Select>
        </FormControl>
    );
};