import Grid from '@mui/material/Grid';
import styled from 'styled-components';

const FooterTab = styled.footer`
    display: flex;
    flex: 1;
    justify-content: center;
    align-items: center;
    position: fixed;
    bottom: 0;
    width: 100%;
    margin-top: 100px;

    a {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-grow: 1;
        color: black;
        text-decoration: none;
    }
      
    p {
        font-size: 1vw;
        margin-right: 10px;
    }

    @media (max-width: 500px) {
        display: flex;
        flex: 1;
        justify-content: center;
        align-items: center;
        position: fixed;
        bottom: 0;
        width: 100%;

      
        a {
            display: flex;
            justify-content: center;
            align-items: center;
            flex-grow: 1;
        }
        
        p {
              font-size: 3vw;
              margin-right: 5px;
        }
      }

`

export default function Footer() {
    return (
    <Grid>
        <FooterTab>
            <a
                href="https://github.com/Nexlson"
                target="_blank"
                rel="noopener noreferrer"
                >
                <p>Built by Nexlson</p> {' '}
                <span className="logo">
                    <img src="/github.ico" alt="Github Logo" width={16} height={16} />
                </span>
            </a>
        </FooterTab>
    </Grid>
    )
}


