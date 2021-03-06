import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

ScrollView {
    id: historyFilterDelegate
    signal loadWallets()
    ListView {
        id: listViewFilters
        
        width: parent.width
        spacing: 10
        
        model: modelFilters
        delegate: HistoryFilterListDelegate {
            property var listAddresses
            width: parent.width
        }

        ModelManager {
            id: modelManager
            
            Component.onCompleted: {
                setWalletManager(walletManager)
            }
        }
        Connections{
            target: historyFilterDelegate
            onLoadWallets:{
                modelFilters.loadModel(walletManager.getWallets())
            }
        }
        
        WalletModel {
            id: modelFilters
                       
            Component.onCompleted: {
                loadModel(walletManager.getWallets())
            }
        }
    }
}