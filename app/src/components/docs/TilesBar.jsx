import styled from 'styled-components';

const InteractiveTile = styled.img`
    width: 60px;
    height: 80px;

    @media (max-width: 500px) {
        width: 30px;
        height: 50px;
    }
    
    :active {
        animation-name: selectTile;
        animation-duration: 0.1s;
    }

    @keyframes selectTile {
        from {transform:translateY(0);}
        to {transform:translateY(-10px);}
    }
`

export function TilesBar(props) {
    return(
        <>
        {
            props.tiles.map(
                (tile, index) => (
                    <InteractiveTile src={"/tiles/"+tile+".jpg"}/>
            ))
        }
        
        </>
    )
}