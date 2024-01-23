package persistence

import (
	"github.com/cossim/coss-server/pkg/utils/time"
	"github.com/cossim/coss-server/service/msg/api/dataTransformers"
	"github.com/cossim/coss-server/service/msg/domain/entity"
	"gorm.io/gorm"
)

type MsgRepo struct {
	db *gorm.DB
}

func NewMsgRepo(db *gorm.DB) *MsgRepo {
	return &MsgRepo{db: db}
}
func (g *MsgRepo) InsertUserMessage(senderId string, receiverId string, msg string, msgType entity.UserMessageType, replyId uint, dialogId uint) (*entity.UserMessage, error) {
	content := &entity.UserMessage{
		SendID:    senderId,
		ReceiveID: receiverId,
		Content:   msg,
		Type:      msgType,
		ReplyId:   replyId,
		DialogId:  dialogId,
	}
	if err := g.db.Save(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func (g *MsgRepo) InsertGroupMessage(uid string, groupId uint, msg string, msgType entity.UserMessageType, replyId uint, dialogId uint) (*entity.GroupMessage, error) {
	content := &entity.GroupMessage{
		UID:      uid,
		GroupID:  groupId,
		Content:  msg,
		Type:     msgType,
		ReplyId:  replyId,
		DialogId: dialogId,
	}
	if err := g.db.Save(content).Error; err != nil {
		return nil, err
	}
	return content, nil
}

func (g *MsgRepo) GetUserMsgList(userId, friendId string, content string, msgType entity.UserMessageType, pageNumber, pageSize int) ([]entity.UserMessage, int32, int32) {
	var results []entity.UserMessage

	query := g.db.Model(&entity.UserMessage{})
	query = query.Where("(send_id = ? AND receive_id = ?) OR (send_id = ? AND receive_id = ?)", userId, friendId, friendId, userId).Order("created_at DESC")

	if content != "" {
		query = query.Where("content LIKE ?", "%"+content+"%")
	}

	if entity.IsValidMessageType(msgType) {
		query = query.Where("msg_type = ?", msgType)
	}

	offset := (pageNumber - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize).Find(&results)
	var count int64
	// 注意：这里的 query 是一个新的查询，需要重新构建条件
	countQuery := g.db.Model(&entity.UserMessage{})
	countQuery = countQuery.Where("(send_id = ? AND receive_id = ?) OR (send_id = ? AND receive_id = ?)", userId, friendId, friendId, userId)
	if content != "" {
		countQuery = countQuery.Where("content LIKE ?", "%"+content+"%")
	}
	if entity.IsValidMessageType(msgType) {
		countQuery = countQuery.Where("msg_type = ?", msgType)
	}
	countQuery.Count(&count)

	return results, int32(count), int32(pageNumber)

}

func (g *MsgRepo) GetLastMsgsForUserWithFriends(userID string, friendIDs []string) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage

	result := g.db.
		Where("(send_id = ? AND receive_id IN (?)) OR (send_id IN (?) AND receive_id = ?)",
			userID, friendIDs, friendIDs, userID).
		Group("receive_id").
		Order("created_at DESC").
		Find(&userMessages)

	if result.Error != nil {
		return nil, result.Error
	}

	return userMessages, nil
}

func (g *MsgRepo) GetLastMsgsForGroupsWithIDs(groupIDs []uint) ([]*entity.GroupMessage, error) {
	var groupMessages []*entity.GroupMessage

	result := g.db.
		Where("group_id IN (?)", groupIDs).
		Group("group_id").
		Order("created_at DESC").
		Find(&groupMessages)

	if result.Error != nil {
		return nil, result.Error
	}

	return groupMessages, nil
}

func (g *MsgRepo) GetLastMsgsByDialogIDs(dialogIds []uint) ([]dataTransformers.LastMessage, error) {
	// 查询 GroupMessage 表中每个 dialog_id 的最后一条数据
	var groupMessages []*entity.GroupMessage
	for _, dialogId := range dialogIds {
		var lastMsg entity.GroupMessage
		g.db.Where("dialog_id =? AND deleted_at = 0", dialogId).Select("id, dialog_id, content, type, uid as send_id, created_at").Order("created_at DESC").First(&lastMsg)
		if lastMsg.ID != 0 {
			groupMessages = append(groupMessages, &lastMsg)
		}
	}
	// 查询 UserMessage 表中每个 dialog_id 的最后一条数据
	var userMessages []*entity.UserMessage
	for _, dialogId := range dialogIds {
		var lastMsg entity.UserMessage
		g.db.Where("dialog_id =? AND deleted_at = 0", dialogId).Select("id, dialog_id, content, type, send_id, created_at").Order("created_at DESC").First(&lastMsg)
		if lastMsg.ID != 0 {
			userMessages = append(userMessages, &lastMsg)
		}
	}

	// 合并两个表的数据
	var result []dataTransformers.LastMessage
	for _, groupMsg := range groupMessages {
		result = append(result, dataTransformers.LastMessage{
			ID:       groupMsg.ID,
			DialogId: groupMsg.DialogId,
			Content:  groupMsg.Content,
			Type:     uint(groupMsg.Type),
			SenderId: groupMsg.UID,
			CreateAt: groupMsg.CreatedAt,
		})
	}
	for _, userMsg := range userMessages {
		result = append(result, dataTransformers.LastMessage{
			ID:       userMsg.ID,
			DialogId: userMsg.DialogId,
			Content:  userMsg.Content,
			Type:     uint(userMsg.Type),
			SenderId: userMsg.SendID,
			CreateAt: userMsg.CreatedAt,
		})
	}
	return result, nil
}

func (g *MsgRepo) UpdateUserMessage(msg entity.UserMessage) error {
	err := g.db.Model(&msg).Updates(msg).Error
	return err
}

func (g *MsgRepo) UpdateGroupMessage(msg entity.GroupMessage) error {
	err := g.db.Model(&msg).Updates(msg).Error
	return err
}

func (g *MsgRepo) LogicalDeleteUserMessage(msgId uint32) error {
	err := g.db.Model(&entity.UserMessage{}).Where("id = ?", msgId).Update("deleted_at", time.Now()).Error
	return err
}

func (g *MsgRepo) LogicalDeleteGroupMessage(msgId uint32) error {
	err := g.db.Model(&entity.GroupMessage{}).Where("id = ?", msgId).Update("deleted_at", time.Now()).Error
	return err
}

func (g *MsgRepo) PhysicalDeleteUserMessage(msgId uint32) error {
	err := g.db.Delete(&entity.UserMessage{}, msgId).Error
	return err
}

func (g *MsgRepo) PhysicalDeleteGroupMessage(msgId uint32) error {
	err := g.db.Delete(&entity.GroupMessage{}, msgId).Error
	return err
}

func (g *MsgRepo) GetUserMsgByID(msgId uint32) (*entity.UserMessage, error) {
	msg := &entity.UserMessage{}
	if err := g.db.Model(&entity.UserMessage{}).Where("id = ? AND deleted_at = 0", msgId).First(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

func (g *MsgRepo) GetGroupMsgByID(msgId uint32) (*entity.GroupMessage, error) {
	msg := &entity.GroupMessage{}
	if err := g.db.Model(&entity.GroupMessage{}).Where("id = ? AND deleted_at = 0", msgId).First(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}

func (g *MsgRepo) UpdateUserMsgColumn(msgId uint32, column string, value interface{}) error {
	return g.db.Model(&entity.UserMessage{}).Where("id = ?", msgId).Update(column, value).Error
}

func (g *MsgRepo) UpdateGroupMsgColumn(msgId uint32, column string, value interface{}) error {
	return g.db.Model(&entity.GroupMessage{}).Where("id = ?", msgId).Update(column, value).Error
}

func (g *MsgRepo) GetUserMsgLabelByDialogId(dialogId uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	if err := g.db.Model(&entity.UserMessage{}).Where("dialog_id = ? AND is_label = ? AND deleted_at = 0", dialogId, entity.IsLabel).Find(&userMessages).Error; err != nil {
		return nil, err
	}
	return userMessages, nil
}

func (g *MsgRepo) GetGroupMsgLabelByDialogId(dialogId uint32) ([]*entity.GroupMessage, error) {
	var groupMessages []*entity.GroupMessage
	if err := g.db.Model(&entity.GroupMessage{}).Where("dialog_id = ? AND is_label = ? AND deleted_at = 0", dialogId, entity.IsLabel).Find(&groupMessages).Error; err != nil {
		return nil, err
	}
	return groupMessages, nil
}

func (g *MsgRepo) SetUserMsgsReadStatus(msgIds []uint32, dialogId uint32) error {
	return g.db.Model(&entity.UserMessage{}).
		Where("id IN (?) AND dialog_id = ? AND deleted_at = 0", msgIds, dialogId).
		Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": time.Now(),
		}).Error
}

func (g *MsgRepo) SetUserMsgReadStatus(msgId uint32, isRead entity.ReadType) error {
	dd := g.db.Model(&entity.UserMessage{}).Where("id = ? AND deleted_at = 0", msgId)
	if isRead == entity.IsRead {
		return dd.Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": time.Now(),
		}).Error
	} else {
		return dd.Updates(map[string]interface{}{
			"is_read": entity.IsRead,
			"read_at": 0,
		}).Error
	}
}

func (g *MsgRepo) GetUnreadUserMsgs(uid string, dialogId uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	if err := g.db.Model(&entity.UserMessage{}).Where("dialog_id = ? AND is_read = ? AND receive_id = ? AND deleted_at = 0", dialogId, entity.NotRead, uid).Find(&userMessages).Error; err != nil {
		return nil, err
	}
	return userMessages, nil
}

func (g *MsgRepo) GetBatchUserMsgsBurnAfterReadingMessages(msgIds []uint32, dialogID uint32) ([]*entity.UserMessage, error) {
	var userMessages []*entity.UserMessage
	err := g.db.Model(&entity.UserMessage{}).
		Where("dialog_id = ? AND id IN (?) AND type = ?", dialogID, msgIds, entity.MessageTypeBurnAfterReading).
		Find(&userMessages).Error
	if err != nil {
		return nil, err
	}
	return userMessages, nil
}

func (g *MsgRepo) PhysicalDeleteUserMessages(msgIds []uint32) error {
	return g.db.Delete(&entity.UserMessage{}, msgIds).Error
}

func (g *MsgRepo) PhysicalDeleteGroupMessages(msgIds []uint32) error {
	return g.db.Delete(&entity.GroupMessage{}, msgIds).Error
}

func (g *MsgRepo) LogicalDeleteUserMessages(msgIds []uint32) error {
	return g.db.Model(&entity.UserMessage{}).
		Where("id IN (?)", msgIds).
		Update("deleted_at", time.Now()).Error
}

func (g *MsgRepo) LogicalDeleteGroupMessages(msgIds []uint32) error {
	return g.db.Model(&entity.GroupMessage{}).
		Where("id IN (?)", msgIds).
		Update("deleted_at", time.Now()).Error
}
