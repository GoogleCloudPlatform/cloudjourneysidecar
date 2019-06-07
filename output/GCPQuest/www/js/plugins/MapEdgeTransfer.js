//=============================================================================
// Map Edge Transfer
// by Shaz
// Last Updated: 2015.10.30
//
// Revisions:
// 2015.10.30 - Fix problem with player being temporarily invisible after transfer
//=============================================================================

/*:
 * @plugindesc Allows auto transfer from edge of map
 * @author Shaz
 *
 * @help This plugin does not provide plugin commands.
 *
 * Add <tfrup: 1 2 3 4 5>
 *     <tfrdown: 1 2 3 4 5>
 *     <tfrleft: 1 2 3 4 5>
 *     <tfrright: 1 2 3 4 5>
 * to map notes to transfer to a new location when the player reaches the
 * edge of the map
 *
 * 1 - map id
 * 2 - x coordinate on new map
 * 3 - y coordinate on new map
 * 4 - facing direction (optional - if omitted, direction is retained)
 * 5 - transfer type (0=black, 1=white, 2=none - optional - if omitted, 0)
 *                   (if entered, facing direction is also required)
 *
 * Each value can be one of:
 * a number
 * $gameVariables.value(n) - to get map, x or y from variable n
 * a formula - to calculate based on a formula (eg: $gamePlayer.y)
 *
 */

(function() {
    var _Game_Player_update = Game_Player.prototype.update;
    Game_Player.prototype.update = function(sceneActive) {
      _Game_Player_update.call(this, sceneActive);
      if (!this.isMoving()) {
        this.mapEdgeTransferUpdate();
      }
    };
  
    Game_Player.prototype.mapEdgeTransferUpdate = function() {
      if (this.x === 0 && this.direction() === 4 && $dataMap.meta.tfrleft) {
        this.mapEdgeTransfer($dataMap.meta.tfrleft);
      }
      if (this.x === $dataMap.width - 1 && this.direction() === 6 && $dataMap.meta.tfrright) {
        this.mapEdgeTransfer($dataMap.meta.tfrright);
      }
      if (this.y === 0 && this.direction() === 8 && $dataMap.meta.tfrup) {
        this.mapEdgeTransfer($dataMap.meta.tfrup);
      }
      if (this.y === $dataMap.height - 1 && this.direction() === 2 && $dataMap.meta.tfrdown) {
        this.mapEdgeTransfer($dataMap.meta.tfrdown);
      }
    };
  
    Game_Player.prototype.mapEdgeTransfer = function(tfrDetails) {
      var tfrInfo = tfrDetails.trim().split(' ');
      if (tfrDetails.length < 3) {
        return;
      };
      var tfrMap = eval(tfrInfo[0]);
      var tfrX = eval(tfrInfo[1]);
      var tfrY = eval(tfrInfo[2]);
      var tfrDir = tfrDetails.length < 4 ? this.direction() : parseInt(eval(tfrInfo[3]));
      if (![2,4,6,8].contains(tfrDir)) {
        tfrDir = this.direction();
      }
      var tfrType = tfrDetails.length < 5 ? 0 : parseInt(eval(tfrInfo[4]));
      if (![0,1,2].contains(tfrType)) {
        tfrType = 0;
      }
      this.reserveTransfer(tfrMap, tfrX, tfrY, tfrDir, tfrType);
      $gameTemp.clearDestination();
    };
  })();