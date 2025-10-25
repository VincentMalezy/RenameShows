# RenameShows

RenameShows est un outil pour renommer automatiquement vos fichiers `.mkv` selon le nom de la série et de la saison.  

## Prérequis

- L’outil ne renomme que les fichiers **.mkv**.  
- Tous les fichiers `.mkv` d’un même dossier doivent appartenir à la **même série** et à la **même saison**.  
- Les fichiers initiaux doivent être au format suivant **avant renommage** :  

Nom.S01E01.XXX.mkv

- `Nom` → nom de la série  
- `S01E01` → saison et numéro d’épisode  
- `XXX` → texte variable ignoré



## Format de renommage

Les fichiers seront renommés selon ce format :  

Episode X - Nom.mkv

- `X` correspond au numéro de l’épisode.  
- `Nom` correspond au titre de l'épisode traduit en français.  

## Utilisation

1. Placez l’exécutable `RenameShows.exe` dans le dossier contenant vos fichiers `.mkv`.  
2. Lancez l’exécutable.  
3. Les fichiers `.mkv` seront renommés automatiquement selon le format ci-dessus.  

> ⚠️ Assurez-vous que le dossier ne contient que les épisodes de la même série et de la même saison pour éviter les erreurs de renommage.