import PhotoOutlined from '@mui/icons-material/PhotoOutlined';
import PlayCircleOutlineOutlined from '@mui/icons-material/PlayCircleOutlineOutlined';
import { styled } from '@mui/material';
import { FILE_TYPE } from 'constants/file';
import { Overlay } from './Container';

interface Iprops {
    fileType: FILE_TYPE;
}

const CenteredOverlay = styled(Overlay)`
    display: flex;
    justify-content: center;
    align-items: center;
`;

export const StaticThumbnail = (props: Iprops) => {
    return (
        <CenteredOverlay
            sx={(theme) => ({
                backgroundColor: theme.palette.fill.dark,
                borderWidth: '1px',
                borderStyle: 'solid',
                borderColor: theme.palette.stroke.faint,
                borderRadius: '4px',
                '& > svg': {
                    color: theme.palette.stroke.muted,
                    fontSize: '50px',
                },
            })}>
            {props.fileType !== FILE_TYPE.VIDEO ? (
                <PhotoOutlined />
            ) : (
                <PlayCircleOutlineOutlined />
            )}
        </CenteredOverlay>
    );
};

export const LoadingThumbnail = () => {
    return (
        <Overlay
            sx={(theme) => ({
                backgroundColor: theme.palette.fill.dark,
                borderWidth: '1px',
                borderStyle: 'solid',
                borderColor: theme.palette.stroke.faint,
                borderRadius: '4px',
            })}
        />
    );
};
